package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetMovementByResourceID(t *testing.T) {
	t.Parallel()

	var date = time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)

	type args struct {
		resourceID types.ID
		filter     movement.Filter
	}
	cases := []struct {
		name    string
		args    args
		want    types.Quantity
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return sum of movements",
			args: args{
				resourceID: 1,
			},

			want: 100,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
					ID: 1,
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				})
				if err != nil {
					t.Error(err)
				}

				_, err = db.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        6,
					ProductID: 1,
					Type:      movement.InputType,
					Date:      date,
					Quantity:  600,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   547556444444444,
				})
				if err != nil {
					t.Error(err)
				}
			},
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)

			a := actions{db: db.Pool}
			got, err := a.GetBalanceByProductID(ctx, tt.args.resourceID, tt.args.filter)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.resourceID)
		})
	}
}
