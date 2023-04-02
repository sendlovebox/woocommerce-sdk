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

	SearchProducts(ctx context.Context, request model.SearchProductsRequest) ([]*model.Product, error)
	GetProductByID(ctx context.Context, id string) (model.Product, error)
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
	//c.SetBaseURL(baseURL)
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
