package item

import (
	"comies/app/core/entities/item"
	"comies/app/core/throw"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, i item.Item) (item.Item, error) {
	const script = `
		insert into items (
			id,
			order_id,
			status,
            price,
			product_id,
			quantity,
			observations
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := transaction.ExecFromContext(ctx, script,
		i.ID,
		i.OrderID,
		i.Status,
		i.Price,
		i.ProductID,
		i.Quantity,
		i.Observations,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.ItemOrderIDFK {
				return item.Item{}, throw.Error(throw.ErrNotFound).
					Describe("the order id provided seems to not exist").Params(map[string]interface{}{
					"order_id": i.OrderID.String(),
				})
			}
			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.ItemIDPK {
				return item.Item{}, throw.Error(throw.ErrAlreadyExists).
					Describe("the item id provided seems to already exist").Params(map[string]interface{}{
					"order_id": i.OrderID.String(), "item_id": i.ID.String(),
				})
			}
		}

		return item.Item{}, throw.Error(err)
	}

	return i, nil
}
