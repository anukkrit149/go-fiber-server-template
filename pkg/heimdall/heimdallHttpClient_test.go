package heimdall

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/hystrix"
	"github.com/stretchr/testify/assert"
)

func getHysterixClientFromConfigTest() heimdall.Doer {
	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(1000*time.Millisecond),
		hystrix.WithHystrixTimeout(1000*time.Millisecond),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
		hystrix.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(60*time.Millisecond, 2*time.Millisecond))),
		hystrix.WithRetryCount(10),
		hystrix.WithFallbackFunc(nil))
}

func TestNewHeimdallClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		switch req.URL.String() {
		case "/":
			rw.Header().Set("Content-Type", "text/plain")
			_, _ = io.WriteString(rw, "This is an example server.")

		default:
			rw.WriteHeader(404)
		}
	}))

	// Close the server when testCase finishes
	defer server.Close()

	client := getHysterixClientFromConfigTest()

	// testing for success at port 8081 as server is running on it
	request, _ := http.NewRequest(
		http.MethodGet,
		server.URL, nil)
	resp, err := client.Do(request)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// testing for failure at 9999
	request, _ = http.NewRequest( // nosemgrep : problem-based-packs.insecure-transport.go-stdlib.http-customized-request.http-customized-request
		http.MethodGet,
		"http://localhost:9999", nil)
	resp, err = client.Do(request)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
