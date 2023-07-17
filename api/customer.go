package api

import (
	"context"
	"net/http"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// CreateACustomer handles logic to create a new customer
func (c *Call) CreateACustomer(ctx context.Context, request model.CreateCustomerRequest) (*model.Customer, error) {
	response := &model.Customer{}

	path := "/customers"
	err := c.makeRequest(ctx, http.MethodPost, path, request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
