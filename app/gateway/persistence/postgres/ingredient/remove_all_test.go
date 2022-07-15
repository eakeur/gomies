package ingredient

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/core/entities/product"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_RemoveAll(t *testing.T) {
	t.Parallel()

	type args struct {
		productID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete ingredients successfully",
			args: args{
				productID: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PZF",
					Name: "Pizza de frango",
					Type: product.OutputType,
				}, product.Product{
					ID:   2,
					Code: "FGO",
					Name: "Frango",
					Type: product.InputType,
				}, product.Product{
					ID:   3,
					Code: "MSS",
					Name: "Massa",
					Type: product.InputType,
				}, product.Product{
					ID:   4,
					Code: "QJO",
					Name: "Queijo",
					Type: product.InputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}

				_, err = d.InsertIngredients(ctx,
					ingredient.Ingredient{ID: 1, ProductID: 1, IngredientID: 2, Quantity: 100},
					ingredient.Ingredient{ID: 2, ProductID: 1, IngredientID: 3, Quantity: 100},
					ingredient.Ingredient{ID: 3, ProductID: 1, IngredientID: 4, Quantity: 100},
				)
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from ingredients", int64(0))
			},
		},
		{
			name: "should fail for nonexistent product",
			args: args{
				productID: 1,
			},
			wantErr: fault.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.RemoveAll(ctx, tt.args.productID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
