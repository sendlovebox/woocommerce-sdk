// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/sendlovebox/woocommerce-sdk/model"
)

// RemoteCalls abstracted definition of supported functions
//
//go:generate mockgen -source api.go -destination ./mock_remote_calls.go -package api RemoteCalls
type RemoteCalls interface {
	Auth(customerKey, customerSecret string) error
	RunInSandboxMode()

	ListAllProducts(ctx context.Context, request model.SearchProductsRequest) ([]*model.Product, error)
	RetrieveAProduct(ctx context.Context, id string) (model.Product, error)
	ListAllProductTags(ctx context.Context, request model.SearchTagsRequest) ([]*model.ProductTag, error)

	RetrieveProductVariations(ctx context.Context, productID string, request model.SearchProductVariationsRequest) ([]model.ProductVariation, error)
	RetrieveAProductVariation(ctx context.Context, productID, variationID string) (model.ProductVariation, error)

	ListAllCategories(ctx context.Context, request model.SearchCategoriesRequest) ([]*model.Category, error)
	RetrieveACategory(ctx context.Context, id string) (model.Category, error)

	CreateACustomer(ctx context.Context, request model.CreateCustomerRequest) (*model.Customer, error)

	CreateAnOrder(ctx context.Context, request model.CreateOrderRequest) (*model.Order, error)
	RetrieveAnOrder(ctx context.Context, orderID int) (*model.Order, error)
	UpdateAnOrder(ctx context.Context, orderID int, request model.UpdateOrderRequest) (*model.Order, error)
	DeleteAnOrder(ctx context.Context, orderID int, forceDelete bool) (*model.Order, error)
}

// Call object
type Call struct {
	client      *resty.Client
	logger      zerolog.Logger
	sandboxMode bool
	baseURL     string
}

// New initialises the object Call
func New(z *zerolog.Logger, c *resty.Client, baseURL string) RemoteCalls {
	c.SetTimeout(15 * time.Second)
	call := &Call{
		client:  c,
		logger:  z.With().Str("sdk", "woocommerce").Logger(),
		baseURL: baseURL,
	}
	return RemoteCalls(call)
}

// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
func (c *Call) RunInSandboxMode() {
	c.client.SetDebug(true)
	c.client.EnableTrace()
	c.sandboxMode = true
}

// Auth to set authorization token
func (c *Call) Auth(customerKey, customerSecret string) error {
	c.client.SetBasicAuth(customerKey, customerSecret)
	return nil
}
