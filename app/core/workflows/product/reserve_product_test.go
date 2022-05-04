package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/types"
	"testing"
)

func TestWorkflow_ReserveProduct(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			reservation Reservation
		}

		opts struct {
			products *product.ActionsMock
			stocks   *StockServiceMock
		}

		test struct {
			name    string
			args    args
			want    ReservationResult
			wantErr error
			opts    opts
		}
	)

	for _, c := range []test{
		{
			name: "should reserve product without ingredients",
			args: args{
				reservation: Reservation{
					ID:        1,
					ProductID: 2,
					Quantity:  10,
					Price:     1000,
				},
			},
			want: ReservationResult{
				Price: 1000,
			},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:   1000,
							MinimumSale: 1,
						}, nil
					},
				},
				stocks: &StockServiceMock{
					ReserveResourcesFunc: func(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error) {
						return nil, nil
					},
				},
			},
		},
		{
			name: "should reserve no ingredient because none is enough",
			args: args{
				reservation: Reservation{
					ID:        1,
					ProductID: 2,
					Quantity:  1000,
					Price:     1000,
				},
			},
			want: ReservationResult{
				Price: 1000,
				FailedChecks: []FailedReservation{
					{
						ProductID: 1,
						Want:      10000,
						Got:       500,
						Error:     product.ErrNotEnoughStocked,
					},
					{
						ProductID: 2,
						Want:      10000,
						Got:       500,
						Error:     product.ErrNotEnoughStocked,
					},
					{
						ProductID: 3,
						Want:      20000,
						Got:       500,
						Error:     product.ErrNotEnoughStocked,
					},
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{
								IngredientID: 1,
								Quantity:     10,
							},
							{
								IngredientID: 2,
								Quantity:     10,
							}, {
								IngredientID: 3,
								Quantity:     20,
							},
						}, nil
					},
				},
				stocks: &StockServiceMock{
					ReserveResourcesFunc: func(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error) {
						return []FailedReservation{
							{
								ProductID: 1,
								Want:      10000,
								Got:       500,
								Error:     product.ErrNotEnoughStocked,
							},
							{
								ProductID: 2,
								Want:      10000,
								Got:       500,
								Error:     product.ErrNotEnoughStocked,
							},
							{
								ProductID: 3,
								Want:      20000,
								Got:       500,
								Error:     product.ErrNotEnoughStocked,
							},
						}, nil
					},
				},
			},
		},
		{
			name: "should reserve some ingredients and fail the ones not stocked",
			args: args{
				reservation: Reservation{
					ID:        1,
					ProductID: 2,
					Quantity:  1000,
					Price:     1000,
				},
			},
			want: ReservationResult{
				Price: 1000,
				FailedChecks: []FailedReservation{
					{
						ProductID: 3,
						Want:      20000,
						Got:       15000,
						Error:     product.ErrNotEnoughStocked,
					},
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{
								IngredientID: 1,
								Quantity:     10,
							},
							{
								IngredientID: 2,
								Quantity:     10,
							}, {
								IngredientID: 3,
								Quantity:     20,
							},
						}, nil
					},
				},
				stocks: &StockServiceMock{
					ReserveResourcesFunc: func(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error) {
						return []FailedReservation{
							{
								ProductID: 3,
								Want:      20000,
								Got:       15000,
								Error:     product.ErrNotEnoughStocked,
							},
						}, nil
					},
				},
			},
		},
		{
			name: "should reserve some ingredients and ignore the ones parameterized",
			args: args{
				reservation: Reservation{
					ID:        1,
					ProductID: 2,
					Quantity:  1000,
					Price:     1000,
					Ignore: []types.ID{
						3,
					},
				},
			},
			want: ReservationResult{
				Price: 1000,
			},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{
								IngredientID: 1,
								Quantity:     10,
							},
							{
								IngredientID: 2,
								Quantity:     10,
							},
							{
								IngredientID: 3,
								Quantity:     20,
							},
						}, nil
					},
				},
				stocks: &StockServiceMock{
					ReserveResourcesFunc: func(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error) {
						for _, resource := range resources {
							if resource.IngredientID == 3 {
								return []FailedReservation{
									{
										ProductID: 3,
										Want:      20000,
										Got:       15000,
										Error:     product.ErrNotEnoughStocked,
									},
								}, nil
							}
						}
						return nil, nil
					},
				},
			},
		},
		{
			name: "should reserve some ingredients and replace the ones parameterized",
			args: args{
				reservation: Reservation{
					ID:        1,
					ProductID: 2,
					Quantity:  1000,
					Price:     1000,
					Replace: map[types.ID]types.ID{
						3: 5,
					},
				},
			},
			want: ReservationResult{
				Price: 1000,
				FailedChecks: []FailedReservation{
					{
						ProductID: 5,
						Want:      20000,
						Got:       15000,
						Error:     product.ErrNotEnoughStocked,
					},
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{
								IngredientID: 1,
								Quantity:     10,
							},
							{
								IngredientID: 2,
								Quantity:     10,
							},
							{
								IngredientID: 3,
								Quantity:     20,
							},
						}, nil
					},
				},
				stocks: &StockServiceMock{
					ReserveResourcesFunc: func(ctx context.Context, reservationID types.ID, resources ...product.Ingredient) ([]FailedReservation, error) {
						for _, resource := range resources {
							if resource.IngredientID == 3 || resource.IngredientID == 5 {
								return []FailedReservation{
									{
										ProductID: resource.IngredientID,
										Want:      20000,
										Got:       15000,
										Error:     product.ErrNotEnoughStocked,
									},
								}, nil
							}
						}
						return nil, nil
					},
				},
			},
		},
	} {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got, gotErr := workflow{products: c.opts.products, stocks: c.opts.stocks}.
				ReserveProduct(context.Background(), c.args.reservation)

			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
