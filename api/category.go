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

// ListAllCategories helps you to view all the products.
func (c *Call) ListAllCategories(ctx context.Context, request model.SearchCategoriesRequest) ([]*model.Category, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, "/products/categories")
	fL := c.logger.With().Str("func", "SearchCategories").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		response   interface{}
		categories []*model.Category
		errorRes   model.ErrorPayload
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

	err = mapstructure.Decode(response, &categories)
	if err != nil {
		fL.Err(err).Msg("unable to decode response for categories")
		return categories, err
	}

	return categories, nil
}

// RetrieveACategory helps you to view a single category by its id
func (c *Call) RetrieveACategory(ctx context.Context, id string) (model.Category, error) {
	endpoint := fmt.Sprintf("%s%s%s", c.baseURL, "/products/categories/", id)
	fL := c.logger.With().Str("func", "RetrieveACategory").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, id).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		response interface{}
		category model.Category
		errorRes model.ErrorPayload
	)

	resp, err := c.client.R().
		SetResult(&response).
		SetError(errorRes).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		return category, errors.New(errorRes.Message)
	} else if resp.StatusCode() != http.StatusOK {
		fL.Info().Str("error_code", fmt.Sprintf("%d", resp.StatusCode())).Msg(string(resp.Body()))
		return category, model.ErrNetworkError
	}

	err = mapstructure.Decode(response, &category)
	if err != nil {
		fL.Err(err).Msg("unable to decode response for product")
		return category, err
	}

	return category, nil
}
