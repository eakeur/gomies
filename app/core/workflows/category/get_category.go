package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/pkg/sdk/fault"
)

func (w workflow) GetCategory(ctx context.Context, key category.Key) (category.Category, error) {
	const operation = "Workflows.Product.GetCategory"

	ct, err := w.categories.GetCategory(ctx, key)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
