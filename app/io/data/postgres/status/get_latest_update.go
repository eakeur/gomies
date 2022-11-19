package status

import (
	"comies/app/core/ordering/status"
	"comies/app/core/types"
	"comies/app/io/data/postgres/conn"
	"context"
)

func (a actions) GetLatestUpdate(ctx context.Context, orderID types.ID) (status.Status, error) {
	const script = `
		select
			s.order_id,
			s.occurred_at,
			s.value
		from
			latest_statuses s
		where
			p.order_id = $1
	`

	row, err := conn.QueryRowFromContext(ctx, script, orderID)
	if err != nil {
		return status.Status{}, err
	}

	var s status.Status
	if err := row.Scan(
		&s.OrderID,
		&s.OccurredAt,
		&s.Value,
	); err != nil {
		// TODO: implement specific error validations
		return status.Status{}, err
	}

	return s, nil
}
