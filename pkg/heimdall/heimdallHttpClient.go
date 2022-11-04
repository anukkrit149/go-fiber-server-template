package heimdall

import (
	"context"
	"go-rest-webserver-template/pkg/errors"
	"net/http"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/hystrix"
)

const (
	errCodeBadRequest      = "client_bad_request_error"
	errHTTPResponse        = "http_response_error"
	errReadingHTTPResponse = "error_reading_http_response"
	errClientRequestFailed = "client_request_failed"
	errInvalidStatusCode   = "invalid_status_code"
)

var (
	// ErrorClientBadRequest error represents request creation failure
	ErrorClientBadRequest = errors.NewClass(errCodeBadRequest, errCodeBadRequest)

	// ErrorHTTPResponse error represents the request failure
	ErrorHTTPResponse = errors.NewClass(errHTTPResponse, errHTTPResponse)

	// ErrorReadingHTTPResponse represents the received response was not of expected format
	ErrorReadingHTTPResponse = errors.NewClass(errReadingHTTPResponse, errReadingHTTPResponse)

	// ErrorClientRequestFailed represents client call failure
	ErrorClientRequestFailed = errors.NewClass(errClientRequestFailed, errClientRequestFailed)

	// ErrorInvalidStatusCode represents the status code is not one which are accepted
	ErrorInvalidStatusCode = errors.NewClass(errInvalidStatusCode, errInvalidStatusCode)
)

// IClient interface for heimdall client
type IClient interface {
	Do(context.Context, *http.Request) (*http.Response, error)
}

type client struct {
	client heimdall.Doer
}

// Do makes the http request with heimdall client
func (hc *client) Do(ctx context.Context, request *http.Request) (*http.Response, error) {

	resp, err := hc.client.Do(request)

	return resp, err
}

// New creates new heimdall client with the given config
func New(config *Config) *client {
	return &client{
		client: getHysterixClientFromConfig(&config.Resiliency, config.Retriable, config.HttpClient),
	}
}

func getHysterixClientFromConfig(
	resiliencyConfig *ResiliencyConfig,
	retrier *RetriableConfig,
	httpClient *http.Client) heimdall.Doer {

	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(time.Duration(resiliencyConfig.CircuitBreakerTimeout)*time.Millisecond),
		hystrix.WithHystrixTimeout(time.Duration(resiliencyConfig.CircuitBreakerTimeout)*time.Millisecond),
		hystrix.WithMaxConcurrentRequests(resiliencyConfig.MaxConcurrentRequests),
		hystrix.WithErrorPercentThreshold(resiliencyConfig.ErrorPercentThreshold),
		hystrix.WithRetrier(retrier.Retrier),
		hystrix.WithRetryCount(retrier.RetryCount),
		hystrix.WithFallbackFunc(retrier.RetryFunc),
		hystrix.WithHTTPClient(httpClient))

	return client
}
