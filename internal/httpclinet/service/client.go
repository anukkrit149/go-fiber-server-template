package service

import (
	"context"
	"gitlab/trell/go-foundation/internal/httpclinet"
	interfaces "gitlab/trell/go-foundation/internal/httpclinet/interfaces"
	"gitlab/trell/go-foundation/internal/providers/logger"
	"gitlab/trell/go-foundation/pkg/errors"
	"gitlab/trell/go-foundation/pkg/heimdall"
	"time"
)

type client struct {
	Host       string
	Auth       httpclinet.Auth
	Timeout    time.Duration
	HttpClient heimdall.IClient
}

type Auth struct {
	Key    string
	Secret string
}

// RequestBody holds the request related data
type RequestBody struct {
	RequestConfig RequestConfig
	Body          interface{}
}

type RequestConfig struct {
	Method string
	Uri    string
}

func NewClient(config *httpclinet.Config) (interfaces.IClient, errors.IError) {

	requestClient := heimdall.New(&config.HttpClient)

	return &client{
		Host:       config.Host,
		Auth:       config.Auth,
		Timeout:    config.Timeout,
		HttpClient: requestClient,
	}, nil

}

func (c *client) MakeRequest(ctx context.Context, req interfaces.IRequest, resp interfaces.IResponse, methodName, serviceName string) errors.IError {
	reqpar, creErr := req.CreateRequest()

	logger.Logger(ctx).Infow("Before Request After making Http Call", map[string]interface{}{
		"Host":       reqpar.Host,
		"Url":        reqpar.URL,
		"RequestURl": reqpar.RequestURI,
		"Method":     reqpar.Method,
	})
	if creErr != nil {
		return creErr
	}

	response, reqErr := c.HttpClient.Do(ctx, reqpar)

	if reqErr != nil {
		return heimdall.ErrorClientRequestFailed.New("").Wrap(reqErr)
	}

	logger.Logger(ctx).Infow("response After making Http Call", map[string]interface{}{
		"responseStatus": response.Status,
		"responseErr":    reqErr,
		"Host":           response.Request.Host,
		"Url":            response.Request.URL,
		"RequestURl":     response.Request.RequestURI,
		"Method":         response.Request.Method,
	})

	if respErr := resp.ReadResponse(response); respErr != nil {
		return respErr
	}

	return nil
}

func (c client) GetHost() string {
	return c.Host
}

// GetAuth gives the basic auth details required to connect to client
func (c client) GetAuth() httpclinet.Auth {
	return c.Auth
}

func (c *client) SetHttpClient(client heimdall.IClient) {
	c.HttpClient = client
}

func (c *client) SetAuth(ctx context.Context, req interfaces.IRequest) {

}
