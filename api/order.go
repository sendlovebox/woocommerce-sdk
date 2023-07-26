package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// CreateAnOrder handles logic to create a new order
//
// Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-an-order
func (c *Call) CreateAnOrder(ctx context.Context, request model.CreateOrderRequest) (*model.Order, error) {
	response := &model.Order{}

	path := "/orders"
	err := c.makeRequest(ctx, http.MethodPost, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveAnOrder handles logic to get an existing order by its ID.
//
// Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#retrieve-an-order
func (c *Call) RetrieveAnOrder(ctx context.Context, orderID int) (*model.Order, error) {
	response := &model.Order{}

	path := fmt.Sprintf("/orders/%d", orderID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAnOrder handles logic to update an existing order by its ID.
//
// Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-an-order
func (c *Call) UpdateAnOrder(ctx context.Context, orderID int, request model.UpdateOrderRequest) (*model.Order, error) {
	response := &model.Order{}

	path := fmt.Sprintf("/orders/%d", orderID)
	err := c.makeRequest(ctx, http.MethodPut, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteAnOrder handles logic to delete an existing order by its ID.
//
// Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#delete-an-order
func (c *Call) DeleteAnOrder(ctx context.Context, orderID int, forceDelete bool) (*model.Order, error) {
	response := &model.Order{}

	request := struct {
		Force bool `json:"force,omitempty"`
	}{
		Force: forceDelete,
	}

	path := fmt.Sprintf("/orders/%d", orderID)
	err := c.makeRequest(ctx, http.MethodDelete, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
