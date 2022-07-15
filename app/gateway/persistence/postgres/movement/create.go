package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/persistence/postgres"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/throw"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func (a actions) Create(ctx context.Context, mov movement.Movement) (movement.Movement, error) {
	const script = `
		insert into movements (
			id,
			stock_id,
			type,
			date,
			quantity,
			value,
			agent_id
		) values (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := transaction.ExecFromContext(ctx, script,
		mov.ID,
		mov.StockID,
		mov.Type,
		mov.Date,
		mov.Quantity,
		mov.PaidValue,
		mov.AgentID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			params := map[string]interface{}{"id": mov.ID, "stock_id": mov.StockID}

			if pgErr.Code == postgres.NonexistentFK && pgErr.ConstraintName == postgres.MovementStockIDFK {
				return movement.Movement{}, throw.Error(throw.ErrNotFound).
					Describe("the stock id provided seems to not exist").Params(params)
			}

			if pgErr.Code == postgres.DuplicateError && pgErr.ConstraintName == postgres.MovementIDPK {
				return movement.Movement{}, throw.Error(throw.ErrAlreadyExists).
					Describe("a movement with this id seems to already exist").Params(params)
			}
		}
		return movement.Movement{}, throw.Error(err)
	}

	return mov, nil
}
