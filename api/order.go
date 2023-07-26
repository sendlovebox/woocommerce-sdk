package api

import (
	"context"
	"net/http"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// CreateAnOrder handles logic to create a new order
func (c *Call) CreateAnOrder(ctx context.Context, request model.CreateOrderRequest) (*model.Order, error) {
	response := &model.Order{}

	path := "/orders"
	err := c.makeRequest(ctx, http.MethodPost, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
