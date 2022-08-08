package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetStockInfoByID(t *testing.T) {
	t.Parallel()

	type args struct {
		productID types.ID
	}
	cases := []struct {
		name    string
		args    args
		want    product.Stock
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return stock data for product id",
			args: args{
				productID: 1,
			},
			want: product.Stock{
				MaximumQuantity: 10,
				MinimumQuantity: 550,
				Location:        "Here",
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDX",
					Name: "Product X",
					Type: product.OutputType,
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 550,
						Location:        "Here",
					},
				}); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrNotFound error for nonexistent product",
			args: args{
				productID: 1,
			},
			wantErr: product.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{db: db.Pool}
			got, err := a.GetStockInfoByID(ctx, tt.args.productID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.productID)
		})
	}
}
