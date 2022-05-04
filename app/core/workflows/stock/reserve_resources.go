package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"sync"
	"time"
)

func (w workflow) ReserveResources(ctx context.Context, reservationID types.ID, reservations []Reservation) ([]ReservationResult, error) {
	const operation = "Workflows.Stock.ReserveResources"

	reservationsNumber := len(reservations)
	responses := make([]ReservationResult, reservationsNumber)
	errChan := make(chan error, reservationsNumber)

	wg := sync.WaitGroup{}
	for i, reservation := range reservations {
		i := i
		reservation := reservation
		wg.Add(1)

		go func() {
			const operation = "Workflows.Stock.ReserveResources.ReservationRoutine"

			defer wg.Done()
			res, err := w.checkResource(ctx, reservationID, reservation)
			if err != nil {
				errChan <- fault.Wrap(err, operation, fault.AdditionalData{
					"reservation_id": reservationID.String(),
					"resource_id":    reservation.ResourceID.String(),
				})
			}
			responses[i] = res
		}()
	}
	wg.Wait()

	if len(errChan) > 0 {
		return nil, fault.Wrap(<-errChan, operation)
	}

	return responses, nil

}

func (w workflow) checkResource(ctx context.Context, reservationID types.ID, reservation Reservation) (ReservationResult, error) {
	const operation = "Workflows.Stock.checkResource"

	movement := stock.Movement{
		ResourceID: reservation.ResourceID,
		Type:       stock.ReservedMovement,
		Date:       time.Now(),
		Quantity:   reservation.Quantity,
		Agent:      reservationID,
	}

	res := ReservationResult{
		ResourceID: reservation.ResourceID,
		Want:       reservation.Quantity,
	}

	actual, err := w.stocks.ComputeStock(ctx, stock.Filter{ResourceID: reservation.ResourceID})
	if err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}
	res.Got = actual

	stk, err := w.stocks.GetStockByID(ctx, reservation.ResourceID)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}

	if err := movement.Validate(); err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}

	actual += movement.Value()

	if actual < stk.MinimumQuantity {
		res.Error = stock.ErrStockEmpty
		return res, nil
	}

	_, err = w.stocks.SaveMovements(ctx, movement)
	if err != nil {
		return ReservationResult{}, fault.Wrap(err, operation)
	}

	return res, nil
}
