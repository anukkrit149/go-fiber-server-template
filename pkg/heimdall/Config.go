package heimdall

import (
	"net/http"

	"github.com/gojektech/heimdall"
)

type Config struct {
	Resiliency ResiliencyConfig
	HttpClient *http.Client
	Retriable  *RetriableConfig
}

type ResiliencyConfig struct {
	MaxConcurrentRequests int
	// RequestVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
	// Default is 20
	RequestVolumeThreshold int
	// CircuitBreakerSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
	// Default is 5000
	CircuitBreakerSleepWindow int
	// ErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests
	// Default is 50
	ErrorPercentThreshold int
	// CircuitBreakerTimeout is how long to wait for command to complete, in milliseconds
	// Default is 1000
	CircuitBreakerTimeout int
	Transport             *http.Transport
}

type RetriableConfig struct {
	Retrier    heimdall.Retriable
	RetryCount int
	RetryFunc  func(error) error
}
