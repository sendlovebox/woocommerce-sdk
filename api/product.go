package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// SearchProducts helps you to view all the products.
func (c *Call) SearchProducts(ctx context.Context, request model.SearchProductsRequest) ([]*model.Product, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/products")
	fL := c.logger.With().Str("func", "SearchProducts").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		response interface{}
		products []*model.Product
		errorRes model.ErrorPayload
	)

	queryParams, err := query.Values(request)
	if err != nil {
		fL.Info().Msg("invalid query params")
		return nil, model.ErrInvalidParams
	}

	resp, err := c.client.R().
		SetResult(&response).
		SetError(errorRes).
		SetQueryParamsFromValues(queryParams).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	err = mapstructure.Decode(response, &products)
	if err != nil {
		fL.Err(err).Msg("unable to decode response for products")
		return products, err
	}

	return products, nil
}

// GetProductByID helps you to view a single product by its id
func (c *Call) GetProductByID(ctx context.Context, id string) (model.Product, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, "/products/", id)
	fL := c.logger.With().Str("func", "GetProductByID").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, id).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		response interface{}
		product  model.Product
		errorRes model.ErrorPayload
	)

	resp, err := c.client.R().
		SetResult(&response).
		SetError(errorRes).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return product, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return product, model.ErrNetworkError
	}

	err = mapstructure.Decode(response, &product)
	if err != nil {
		fL.Err(err).Msg("unable to decode response for product")
		return product, err
	}

	return product, nil
}
