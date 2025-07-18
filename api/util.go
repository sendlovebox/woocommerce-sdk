package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// makeRequest is a super function that handles all requests to the woocommerce SDK
func (c *Call) makeRequest(ctx context.Context, method, path string, body, params, expectedRes interface{}) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")
	log.Info().Interface(model.LogStrRequest, body).Msg("request")
	defer log.Info().Msg("done...")

	var (
		err         error
		res         *resty.Response
		queryParams url.Values
		errorRes    model.ErrorPayload
	)

	client := c.client.R().
		SetContext(ctx).
		SetError(&errorRes)

	// set the query params if it is passed
	if params != nil {
		queryParams, err = query.Values(params)
		if err != nil {
			log.Info().Msg("invalid query params")
			return model.ErrInvalidParams
		}
		client.SetQueryParamsFromValues(queryParams)
	}

	if body != nil {
		client.SetBody(body)
	}

	switch method {
	case http.MethodGet:
		res, err = client.Get(endpoint)
	case http.MethodPost:
		res, err = client.Post(endpoint)
	case http.MethodPut:
		res, err = client.Put(endpoint)
	case http.MethodPatch:
		res, err = client.Patch(endpoint)
	case http.MethodDelete:
		res, err = client.Delete(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return err
	}

	if err != nil {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return err
	}

	if errorRes.Message != "" {
		log.Err(errors.New(errorRes.Message)).Msg("error returned from WooCommerce")
		return errors.New(errorRes.Message)
	}

	// Handle and decode response manually
	cleaned := removeBOM(res.Body())

	var decoded interface{}
	if err = json.Unmarshal(cleaned, &decoded); err != nil {
		log.Err(err).Msg("failed to unmarshal cleaned response body")
		return err
	}

	if err = mapstruct(decoded, expectedRes); err != nil {
		log.Err(err).Msg("failed to mapstruct response")
		return err
	}

	log.Info().Interface(model.LogStrResponse, expectedRes).Msg("response")
	return nil
}

func (c *Call) makePaginatedRequest(ctx context.Context, method, path string, body, params, expectedRes interface{}, count, pageNo int) (model.PageInfo, error) {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")
	log.Info().Interface(model.LogStrRequest, body).Msg("request")
	defer log.Info().Msg("done...")

	var (
		err         error
		res         *resty.Response
		queryParams url.Values
		errorRes    model.ErrorPayload
		pageInfo    model.PageInfo
	)

	client := c.client.R().
		SetContext(ctx).
		SetError(&errorRes)

	// set the query params if it is passed
	if params != nil {
		queryParams, err = query.Values(params)
		if err != nil {
			log.Info().Msg("invalid query params")
			return pageInfo, model.ErrInvalidParams
		}

		client.SetQueryParamsFromValues(queryParams)
	}

	if body != nil {
		client.SetBody(body)
	}

	switch method {
	case http.MethodGet:
		res, err = client.Get(endpoint)
	case http.MethodPost:
		res, err = client.Post(endpoint)
	case http.MethodPut:
		res, err = client.Put(endpoint)
	case http.MethodPatch:
		res, err = client.Patch(endpoint)
	case http.MethodDelete:
		res, err = client.Delete(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return pageInfo, err
	}

	if err != nil {
		log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		return pageInfo, err
	}

	// check to ensure error is not empty
	if errorRes != (model.ErrorPayload{}) {
		err = errors.New(errorRes.Message)
		log.Err(err).Msg("error while making request")
		return pageInfo, err
	}

	totalCount, _ := strconv.Atoi(res.Header().Get("X-Wp-Total"))
	totalPages, _ := strconv.Atoi(res.Header().Get("X-Wp-Totalpages"))
	pageInfo = model.PageInfo{
		Page:            pageNo,
		Size:            count,
		HasNextPage:     pageNo < totalPages,
		HasPreviousPage: pageNo > 1,
		TotalCount:      int64(totalCount),
	}
	cleaned := removeBOM(res.Body())

	// decode cleaned body into interface{}
	var decoded interface{}
	if err = json.Unmarshal(cleaned, &decoded); err != nil {
		log.Err(err).Msg("failed to unmarshal cleaned body")
		return pageInfo, err
	}

	return pageInfo, mapstruct(decoded, expectedRes)
}

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

func removeBOM(data []byte) []byte {
	bom := []byte{0xEF, 0xBB, 0xBF}
	if bytes.HasPrefix(data, bom) {
		return data[len(bom):]
	}
	return data
}
