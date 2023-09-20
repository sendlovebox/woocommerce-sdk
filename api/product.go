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

// ListAllProducts helps you to view all the products.
func (c *Call) ListAllProducts(ctx context.Context, request model.SearchProductsRequest) ([]model.Product, error) {
	response := &[]model.Product{}

	err := c.makeRequest(ctx, http.MethodGet, "/products", nil, request, response)
	if err != nil {
		return nil, err
	}

	return *response, nil
}

// RetrieveAProduct helps you to view a single product by its id
func (c *Call) RetrieveAProduct(ctx context.Context, id string) (model.Product, error) {
	response := &model.Product{}

	path := fmt.Sprintf("%s%s", "/products/", id)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, nil, response)
	if err != nil {
		return model.Product{}, err
	}

	return *response, nil
}

// RetrieveProductVariations helps you to view product variations by a single product id
func (c *Call) RetrieveProductVariations(ctx context.Context, productID string, request model.SearchProductVariationsRequest) ([]model.ProductVariation, error) {
	response := &[]model.ProductVariation{}

	path := fmt.Sprintf("/products/%s/variations", productID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, request, response)
	if err != nil {
		return nil, err
	}

	return *response, nil
}

// RetrieveAProductVariation helps you to view a product variations by its id
func (c *Call) RetrieveAProductVariation(ctx context.Context, productID, variationID string) (model.ProductVariation, error) {
	response := &model.ProductVariation{}

	path := fmt.Sprintf("/products/%s/variations/%s", productID, variationID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, nil, response)
	if err != nil {
		return model.ProductVariation{}, err
	}

	return *response, nil
}

// ListAllProductTags helps you to view all the product tags.
func (c *Call) ListAllProductTags(ctx context.Context, request model.SearchTagsRequest) ([]*model.ProductTag, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/products/tags")
	fL := c.logger.With().Str("func", "ListAllProductTags").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		response interface{}
		tags     []*model.ProductTag
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
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return nil, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return nil, model.ErrNetworkError
	}

	err = mapstructure.Decode(response, &tags)
	if err != nil {
		fL.Err(err).Msg("unable to decode response for products")
		return tags, err
	}

	return tags, nil
}
