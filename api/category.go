package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// ListAllCategories fetches all product categories
func (c *Call) ListAllCategories(ctx context.Context, request model.SearchCategoriesRequest) ([]*model.Category, error) {
	endpoint := fmt.Sprintf("%s/products/categories", c.baseURL)
	fL := c.logger.With().Str("func", "SearchCategories").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Interface(model.LogStrRequest, request).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		categories []*model.Category
		errorRes   model.ErrorPayload
	)

	queryParams, err := query.Values(request)
	if err != nil {
		fL.Info().Msg("invalid query params")
		return nil, model.ErrInvalidParams
	}

	resp, err := c.client.R().
		SetError(&errorRes).
		SetQueryParamsFromValues(queryParams).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("request error")
		return nil, err
	}

	if errorRes.Message != "" {
		fL.Err(err).Msg("api returned error")
		return nil, errors.New(errorRes.Message)
	}

	body := removeBOM(resp.Body())

	var decoded interface{}
	if err = json.Unmarshal(body, &decoded); err != nil {
		fL.Err(err).Msg("failed to unmarshal cleaned response body")
		return nil, err
	}

	if err = mapstructure.Decode(decoded, &categories); err != nil {
		fL.Err(err).Msg("unable to decode response for categories")
		return nil, err
	}

	return categories, nil
}

// RetrieveACategory fetches a single category by its ID
func (c *Call) RetrieveACategory(ctx context.Context, id string) (model.Category, error) {
	endpoint := fmt.Sprintf("%s/products/categories/%s", c.baseURL, id)
	fL := c.logger.With().Str("func", "RetrieveACategory").Str(model.LogStrKeyEndpoint, endpoint).Logger()
	fL.Info().Msg("starting...")
	fL.Info().Str("category_id", id).Msg("request")
	defer fL.Info().Msg("done...")

	var (
		category model.Category
		errorRes model.ErrorPayload
	)

	resp, err := c.client.R().
		SetError(&errorRes).
		SetContext(ctx).
		Get(endpoint)

	if err != nil {
		fL.Err(err).Msg("request error")
		return category, err
	}

	if errorRes.Message != "" {
		fL.Err(err).Msg("api returned error")
		return category, errors.New(errorRes.Message)
	}

	body := removeBOM(resp.Body())

	var decoded interface{}
	if err = json.Unmarshal(body, &decoded); err != nil {
		fL.Err(err).Msg("failed to unmarshal cleaned response body")
		return category, err
	}

	if err = mapstructure.Decode(decoded, &category); err != nil {
		fL.Err(err).Msg("unable to decode response for category")
		return category, err
	}

	return category, nil
}
