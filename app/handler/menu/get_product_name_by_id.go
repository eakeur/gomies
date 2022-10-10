package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/handler/rest"
	"context"
	"net/http"
)

// GetProductNameByID fetches a product name by its ID or code.
//
// @Summary     Fetches a product
// @Description Fetches a product name by its id.
// @Tags        Product
// @Param       product_key path     string false "The product ID"
// @Success     200         {object} rest.Response{data=GetProductNameResponse{}}
// @Failure     404         {object} rest.Response{error=rest.Error{}} "PRODUCT_NOT_FOUND"
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /menu/products/{product_id}/name [GET]
func GetProductNameByID(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "product_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	name, err := menu.GetProductNameByID(ctx, id)
	if err != nil {
		return rest.Response{}
	}
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusOK, GetProductNameResponse{Name: name})
}

type GetProductNameResponse struct {
	// Name is the official name of the product, not exactly the name that the customer sees, but indeed the name
	// shown in fiscal documents
	Name string `json:"name"`
}
