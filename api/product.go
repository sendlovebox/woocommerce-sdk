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
func (c *Call) ListAllProducts(ctx context.Context, request model.SearchProductsRequest) ([]model.Product, model.PageInfo, error) {
	response := &[]model.Product{}

	pageInfo, err := c.makePaginatedRequest(ctx, http.MethodGet, "/products", nil, request, response, request.PerPage, request.Page)
	if err != nil {
		return nil, model.PageInfo{}, err
	}

	return *response, pageInfo, nil
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

// CreateProduct helps to create a product. https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product
func (c *Call) CreateProduct(ctx context.Context, request model.ProductCreateRequest) (*model.Product, error) {
	response := &model.Product{}

	path := "/products"
	err := c.makeRequest(ctx, http.MethodPost, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateProduct helps you to update an existing product. https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-a-product
func (c *Call) UpdateProduct(ctx context.Context, id string, request model.ProductUpdateRequest) (*model.Product, error) {
	response := &model.Product{}

	path := fmt.Sprintf("/products/%s", id)

	err := c.makeRequest(ctx, http.MethodPut, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReviewProduct helps you to create a new product review. https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#create-a-product-review
func (c *Call) ReviewProduct(ctx context.Context, request model.ReviewProductRequest) (*model.ProductReview, error) {
	response := &model.ProductReview{}

	err := c.makeRequest(ctx, http.MethodPost, "/products/reviews", request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateProductReview helps you to update an existing product review. https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#update-a-product-review
func (c *Call) UpdateProductReview(ctx context.Context, id string, request model.ReviewProductRequest) (*model.ProductReview, error) {
	response := &model.ProductReview{}

	path := fmt.Sprintf("/products/reviews/%s", id)

	err := c.makeRequest(ctx, http.MethodPut, path, request, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListProductReviews retrieves all product reviews. https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#list-all-product-reviews
func (c *Call) ListProductReviews(ctx context.Context, request model.SearchProductReviewsRequest) ([]model.ProductReview, error) {
	response := &[]model.ProductReview{}

	err := c.makeRequest(ctx, http.MethodGet, "/products/reviews", nil, request, response)
	if err != nil {
		return nil, err
	}

	return *response, nil
}

// GetAProductReview retrieve a single product review by ID. https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#create-a-product-review
func (c *Call) GetAProductReview(ctx context.Context, reviewID string) (*model.ProductReview, error) {
	response := &model.ProductReview{}

	path := fmt.Sprintf("/products/reviews/%s", reviewID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListAllProductAttributes helps to retrieve all the product attributes
func (c *Call) ListAllProductAttributes(ctx context.Context, request model.SearchProductAttributesRequest) ([]model.ProductAttributeSearchResponse, error) {
	response := &[]model.ProductAttributeSearchResponse{}

	err := c.makeRequest(ctx, http.MethodGet, "/products/attributes", nil, request, response)
	if err != nil {
		return nil, err
	}

	return *response, nil
}

// ListAllProductAttributeTerms helps to retrieve all the product attributes terms. https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-attribute-terms
func (c *Call) ListAllProductAttributeTerms(ctx context.Context, attributeID int, request model.SearchProductAttributeTermRequest) ([]model.ProductAttributeTermSearchResponse, error) {
	response := &[]model.ProductAttributeTermSearchResponse{}

	path := fmt.Sprintf("/products/attributes/%d/terms", attributeID)
	err := c.makeRequest(ctx, http.MethodGet, path, nil, request, response)
	if err != nil {
		return nil, err
	}

	return *response, nil
}
