package service

import (
	"bytes"
	"context"
	"encoding/json"
	goErr "errors"
	"gitlab/trell/go-foundation/internal/config"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab/trell/go-foundation/internal/httpclinet/interfaces"
	"gitlab/trell/go-foundation/pkg/errors"
	"gitlab/trell/go-foundation/pkg/heimdall"
)

const (
	DefaultContentType = "application/json"
	DefaultTimeout     = 30 * time.Second
)

type Base struct {
	Ctx    context.Context
	Client interfaces.IClient
	Auth   config.Auth
}

// RequestConfig request specific configurations
func (b Base) CreateRequest(r RequestBody) (*http.Request, errors.IError) {
	bytesArray, err := json.Marshal(r.Body)
	if err != nil {
		return nil, heimdall.ErrorClientBadRequest.New("").Wrap(err)
	}

	var req *http.Request

	req, err = http.NewRequest(
		r.RequestConfig.Method,
		b.Client.GetHost()+r.RequestConfig.Uri,
		bytes.NewReader(bytesArray))

	if err != nil {
		return nil, heimdall.ErrorClientBadRequest.New("").Wrap(err)
	}

	req.SetBasicAuth(b.Client.GetAuth().Key, b.Client.GetAuth().Secret)

	req.Header.Set("Content-Type", DefaultContentType)

	return req, nil
}
func (b Base) SetAuthHeaders(req *http.Request) (*http.Request, errors.IError) {
	req.SetBasicAuth(b.Auth.Username, b.Auth.Password)
	return req, nil
}

func (b Base) SetHeader(req *http.Request, key string, value string) (*http.Request, errors.IError) {
	req.Header.Set(key, value)
	return req, nil
}

func (b Base) ReadResponse(resp *http.Response, v interface{}) errors.IError {
	if resp == nil {
		return heimdall.ErrorHTTPResponse.New("").Wrap(goErr.New("api http response error"))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return heimdall.ErrorReadingHTTPResponse.New("").Wrap(err)
	}

	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		return heimdall.ErrorReadingHTTPResponse.New("").Wrap(err)
	}

	return nil
}
