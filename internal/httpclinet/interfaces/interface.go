package interfaces

import (
	"context"
	"go-rest-webserver-template/internal/httpclinet"
	"go-rest-webserver-template/pkg/errors"
	"go-rest-webserver-template/pkg/heimdall"
	"net/http"
)

type IClient interface {
	MakeRequest(ctx context.Context, req IRequest, resp IResponse, methodName, serviceName string) errors.IError
	GetHost() string
	GetAuth() httpclinet.Auth
	SetAuth(ctx context.Context, req IRequest)
	SetHttpClient(client heimdall.IClient)
}
type Request struct {
	req *http.Request
}

type Response struct {
	req *http.Request
}

type IRequest interface {
	CreateRequest() (*http.Request, errors.IError)
}
type IResponse interface {
	ReadResponse(resp *http.Response) errors.IError
}
