// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/types"
	"sync"
)

// Ensure, that WorkflowMock does implement Workflow.
// If this is not the case, regenerate this file with moq.
var _ Workflow = &WorkflowMock{}

// WorkflowMock is a mock implementation of Workflow.
//
// 	func TestSomethingThatUsesWorkflow(t *testing.T) {
//
// 		// make and configure a mocked Workflow
// 		mockedWorkflow := &WorkflowMock{
// 			CreateIngredientFunc: func(ctx context.Context, productKey product.Key, input IngredientInput) (product.Ingredient, error) {
// 				panic("mock out the CreateIngredient method")
// 			},
// 			CreateProductFunc: func(ctx context.Context, prd product.Product) (product.Product, error) {
// 				panic("mock out the CreateProduct method")
// 			},
// 			GetProductFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
// 				panic("mock out the GetProduct method")
// 			},
// 			ListProductsFunc: func(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error) {
// 				panic("mock out the ListProducts method")
// 			},
// 			RemoveIngredientFunc: func(ctx context.Context, productKey product.Key, id types.ID) error {
// 				panic("mock out the RemoveIngredient method")
// 			},
// 			RemoveProductFunc: func(ctx context.Context, key product.Key) error {
// 				panic("mock out the RemoveProduct method")
// 			},
// 			ReserveProductFunc: func(ctx context.Context, reservation Reservation) (ReservationResult, error) {
// 				panic("mock out the ReserveProduct method")
// 			},
// 			UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
// 				panic("mock out the UpdateProduct method")
// 			},
// 			UpdateReservationFunc: func(ctx context.Context, reservationID types.ID, consume bool) error {
// 				panic("mock out the UpdateReservation method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// CreateIngredientFunc mocks the CreateIngredient method.
	CreateIngredientFunc func(ctx context.Context, productKey product.Key, input IngredientInput) (product.Ingredient, error)

	// CreateProductFunc mocks the CreateProduct method.
	CreateProductFunc func(ctx context.Context, prd product.Product) (product.Product, error)

	// GetProductFunc mocks the GetProduct method.
	GetProductFunc func(ctx context.Context, key product.Key) (product.Product, error)

	// ListProductsFunc mocks the ListProducts method.
	ListProductsFunc func(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error)

	// RemoveIngredientFunc mocks the RemoveIngredient method.
	RemoveIngredientFunc func(ctx context.Context, productKey product.Key, id types.ID) error

	// RemoveProductFunc mocks the RemoveProduct method.
	RemoveProductFunc func(ctx context.Context, key product.Key) error

	// ReserveProductFunc mocks the ReserveProduct method.
	ReserveProductFunc func(ctx context.Context, reservation Reservation) (ReservationResult, error)

	// UpdateProductFunc mocks the UpdateProduct method.
	UpdateProductFunc func(ctx context.Context, prd product.Product) error

	// UpdateReservationFunc mocks the UpdateReservation method.
	UpdateReservationFunc func(ctx context.Context, reservationID types.ID, consume bool) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateIngredient holds details about calls to the CreateIngredient method.
		CreateIngredient []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey product.Key
			// Input is the input argument value.
			Input IngredientInput
		}
		// CreateProduct holds details about calls to the CreateProduct method.
		CreateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd product.Product
		}
		// GetProduct holds details about calls to the GetProduct method.
		GetProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key product.Key
		}
		// ListProducts holds details about calls to the ListProducts method.
		ListProducts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductFilter is the productFilter argument value.
			ProductFilter product.Filter
		}
		// RemoveIngredient holds details about calls to the RemoveIngredient method.
		RemoveIngredient []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey product.Key
			// ID is the id argument value.
			ID types.ID
		}
		// RemoveProduct holds details about calls to the RemoveProduct method.
		RemoveProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key product.Key
		}
		// ReserveProduct holds details about calls to the ReserveProduct method.
		ReserveProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Reservation is the reservation argument value.
			Reservation Reservation
		}
		// UpdateProduct holds details about calls to the UpdateProduct method.
		UpdateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd product.Product
		}
		// UpdateReservation holds details about calls to the UpdateReservation method.
		UpdateReservation []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReservationID is the reservationID argument value.
			ReservationID types.ID
			// Consume is the consume argument value.
			Consume bool
		}
	}
	lockCreateIngredient  sync.RWMutex
	lockCreateProduct     sync.RWMutex
	lockGetProduct        sync.RWMutex
	lockListProducts      sync.RWMutex
	lockRemoveIngredient  sync.RWMutex
	lockRemoveProduct     sync.RWMutex
	lockReserveProduct    sync.RWMutex
	lockUpdateProduct     sync.RWMutex
	lockUpdateReservation sync.RWMutex
}

// CreateIngredient calls CreateIngredientFunc.
func (mock *WorkflowMock) CreateIngredient(ctx context.Context, productKey product.Key, input IngredientInput) (product.Ingredient, error) {
	if mock.CreateIngredientFunc == nil {
		panic("WorkflowMock.CreateIngredientFunc: method is nil but Workflow.CreateIngredient was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey product.Key
		Input      IngredientInput
	}{
		Ctx:        ctx,
		ProductKey: productKey,
		Input:      input,
	}
	mock.lockCreateIngredient.Lock()
	mock.calls.CreateIngredient = append(mock.calls.CreateIngredient, callInfo)
	mock.lockCreateIngredient.Unlock()
	return mock.CreateIngredientFunc(ctx, productKey, input)
}

// CreateIngredientCalls gets all the calls that were made to CreateIngredient.
// Check the length with:
//     len(mockedWorkflow.CreateIngredientCalls())
func (mock *WorkflowMock) CreateIngredientCalls() []struct {
	Ctx        context.Context
	ProductKey product.Key
	Input      IngredientInput
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey product.Key
		Input      IngredientInput
	}
	mock.lockCreateIngredient.RLock()
	calls = mock.calls.CreateIngredient
	mock.lockCreateIngredient.RUnlock()
	return calls
}

// CreateProduct calls CreateProductFunc.
func (mock *WorkflowMock) CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {
	if mock.CreateProductFunc == nil {
		panic("WorkflowMock.CreateProductFunc: method is nil but Workflow.CreateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd product.Product
	}{
		Ctx: ctx,
		Prd: prd,
	}
	mock.lockCreateProduct.Lock()
	mock.calls.CreateProduct = append(mock.calls.CreateProduct, callInfo)
	mock.lockCreateProduct.Unlock()
	return mock.CreateProductFunc(ctx, prd)
}

// CreateProductCalls gets all the calls that were made to CreateProduct.
// Check the length with:
//     len(mockedWorkflow.CreateProductCalls())
func (mock *WorkflowMock) CreateProductCalls() []struct {
	Ctx context.Context
	Prd product.Product
} {
	var calls []struct {
		Ctx context.Context
		Prd product.Product
	}
	mock.lockCreateProduct.RLock()
	calls = mock.calls.CreateProduct
	mock.lockCreateProduct.RUnlock()
	return calls
}

// GetProduct calls GetProductFunc.
func (mock *WorkflowMock) GetProduct(ctx context.Context, key product.Key) (product.Product, error) {
	if mock.GetProductFunc == nil {
		panic("WorkflowMock.GetProductFunc: method is nil but Workflow.GetProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key product.Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetProduct.Lock()
	mock.calls.GetProduct = append(mock.calls.GetProduct, callInfo)
	mock.lockGetProduct.Unlock()
	return mock.GetProductFunc(ctx, key)
}

// GetProductCalls gets all the calls that were made to GetProduct.
// Check the length with:
//     len(mockedWorkflow.GetProductCalls())
func (mock *WorkflowMock) GetProductCalls() []struct {
	Ctx context.Context
	Key product.Key
} {
	var calls []struct {
		Ctx context.Context
		Key product.Key
	}
	mock.lockGetProduct.RLock()
	calls = mock.calls.GetProduct
	mock.lockGetProduct.RUnlock()
	return calls
}

// ListProducts calls ListProductsFunc.
func (mock *WorkflowMock) ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error) {
	if mock.ListProductsFunc == nil {
		panic("WorkflowMock.ListProductsFunc: method is nil but Workflow.ListProducts was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ProductFilter product.Filter
	}{
		Ctx:           ctx,
		ProductFilter: productFilter,
	}
	mock.lockListProducts.Lock()
	mock.calls.ListProducts = append(mock.calls.ListProducts, callInfo)
	mock.lockListProducts.Unlock()
	return mock.ListProductsFunc(ctx, productFilter)
}

// ListProductsCalls gets all the calls that were made to ListProducts.
// Check the length with:
//     len(mockedWorkflow.ListProductsCalls())
func (mock *WorkflowMock) ListProductsCalls() []struct {
	Ctx           context.Context
	ProductFilter product.Filter
} {
	var calls []struct {
		Ctx           context.Context
		ProductFilter product.Filter
	}
	mock.lockListProducts.RLock()
	calls = mock.calls.ListProducts
	mock.lockListProducts.RUnlock()
	return calls
}

// RemoveIngredient calls RemoveIngredientFunc.
func (mock *WorkflowMock) RemoveIngredient(ctx context.Context, productKey product.Key, id types.ID) error {
	if mock.RemoveIngredientFunc == nil {
		panic("WorkflowMock.RemoveIngredientFunc: method is nil but Workflow.RemoveIngredient was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey product.Key
		ID         types.ID
	}{
		Ctx:        ctx,
		ProductKey: productKey,
		ID:         id,
	}
	mock.lockRemoveIngredient.Lock()
	mock.calls.RemoveIngredient = append(mock.calls.RemoveIngredient, callInfo)
	mock.lockRemoveIngredient.Unlock()
	return mock.RemoveIngredientFunc(ctx, productKey, id)
}

// RemoveIngredientCalls gets all the calls that were made to RemoveIngredient.
// Check the length with:
//     len(mockedWorkflow.RemoveIngredientCalls())
func (mock *WorkflowMock) RemoveIngredientCalls() []struct {
	Ctx        context.Context
	ProductKey product.Key
	ID         types.ID
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey product.Key
		ID         types.ID
	}
	mock.lockRemoveIngredient.RLock()
	calls = mock.calls.RemoveIngredient
	mock.lockRemoveIngredient.RUnlock()
	return calls
}

// RemoveProduct calls RemoveProductFunc.
func (mock *WorkflowMock) RemoveProduct(ctx context.Context, key product.Key) error {
	if mock.RemoveProductFunc == nil {
		panic("WorkflowMock.RemoveProductFunc: method is nil but Workflow.RemoveProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key product.Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockRemoveProduct.Lock()
	mock.calls.RemoveProduct = append(mock.calls.RemoveProduct, callInfo)
	mock.lockRemoveProduct.Unlock()
	return mock.RemoveProductFunc(ctx, key)
}

// RemoveProductCalls gets all the calls that were made to RemoveProduct.
// Check the length with:
//     len(mockedWorkflow.RemoveProductCalls())
func (mock *WorkflowMock) RemoveProductCalls() []struct {
	Ctx context.Context
	Key product.Key
} {
	var calls []struct {
		Ctx context.Context
		Key product.Key
	}
	mock.lockRemoveProduct.RLock()
	calls = mock.calls.RemoveProduct
	mock.lockRemoveProduct.RUnlock()
	return calls
}

// ReserveProduct calls ReserveProductFunc.
func (mock *WorkflowMock) ReserveProduct(ctx context.Context, reservation Reservation) (ReservationResult, error) {
	if mock.ReserveProductFunc == nil {
		panic("WorkflowMock.ReserveProductFunc: method is nil but Workflow.ReserveProduct was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Reservation Reservation
	}{
		Ctx:         ctx,
		Reservation: reservation,
	}
	mock.lockReserveProduct.Lock()
	mock.calls.ReserveProduct = append(mock.calls.ReserveProduct, callInfo)
	mock.lockReserveProduct.Unlock()
	return mock.ReserveProductFunc(ctx, reservation)
}

// ReserveProductCalls gets all the calls that were made to ReserveProduct.
// Check the length with:
//     len(mockedWorkflow.ReserveProductCalls())
func (mock *WorkflowMock) ReserveProductCalls() []struct {
	Ctx         context.Context
	Reservation Reservation
} {
	var calls []struct {
		Ctx         context.Context
		Reservation Reservation
	}
	mock.lockReserveProduct.RLock()
	calls = mock.calls.ReserveProduct
	mock.lockReserveProduct.RUnlock()
	return calls
}

// UpdateProduct calls UpdateProductFunc.
func (mock *WorkflowMock) UpdateProduct(ctx context.Context, prd product.Product) error {
	if mock.UpdateProductFunc == nil {
		panic("WorkflowMock.UpdateProductFunc: method is nil but Workflow.UpdateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd product.Product
	}{
		Ctx: ctx,
		Prd: prd,
	}
	mock.lockUpdateProduct.Lock()
	mock.calls.UpdateProduct = append(mock.calls.UpdateProduct, callInfo)
	mock.lockUpdateProduct.Unlock()
	return mock.UpdateProductFunc(ctx, prd)
}

// UpdateProductCalls gets all the calls that were made to UpdateProduct.
// Check the length with:
//     len(mockedWorkflow.UpdateProductCalls())
func (mock *WorkflowMock) UpdateProductCalls() []struct {
	Ctx context.Context
	Prd product.Product
} {
	var calls []struct {
		Ctx context.Context
		Prd product.Product
	}
	mock.lockUpdateProduct.RLock()
	calls = mock.calls.UpdateProduct
	mock.lockUpdateProduct.RUnlock()
	return calls
}

// UpdateReservation calls UpdateReservationFunc.
func (mock *WorkflowMock) UpdateReservation(ctx context.Context, reservationID types.ID, consume bool) error {
	if mock.UpdateReservationFunc == nil {
		panic("WorkflowMock.UpdateReservationFunc: method is nil but Workflow.UpdateReservation was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ReservationID types.ID
		Consume       bool
	}{
		Ctx:           ctx,
		ReservationID: reservationID,
		Consume:       consume,
	}
	mock.lockUpdateReservation.Lock()
	mock.calls.UpdateReservation = append(mock.calls.UpdateReservation, callInfo)
	mock.lockUpdateReservation.Unlock()
	return mock.UpdateReservationFunc(ctx, reservationID, consume)
}

// UpdateReservationCalls gets all the calls that were made to UpdateReservation.
// Check the length with:
//     len(mockedWorkflow.UpdateReservationCalls())
func (mock *WorkflowMock) UpdateReservationCalls() []struct {
	Ctx           context.Context
	ReservationID types.ID
	Consume       bool
} {
	var calls []struct {
		Ctx           context.Context
		ReservationID types.ID
		Consume       bool
	}
	mock.lockUpdateReservation.RLock()
	calls = mock.calls.UpdateReservation
	mock.lockUpdateReservation.RUnlock()
	return calls
}
