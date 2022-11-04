package httpclinet

import (
	"gitlab/trell/go-foundation/pkg/heimdall"
	"time"
)

type Config struct {
	// Endpoint: reminder url : for example : https://api.trell.co/v1
	Host string

	// Auth: basic auth consisting of key and secret
	Auth Auth

	// Timeout: request timeout config in seconds
	Timeout time.Duration

	// HttpClient: configuration for HTTP client.
	HttpClient heimdall.Config

	// HttpClientName: name for registering the HTTP client.
	HttpClientName string
}

type Auth struct {
	Key    string
	Secret string
}
