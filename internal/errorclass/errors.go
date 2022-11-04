package errorclass

import (
	"go-rest-webserver-template/pkg/errors"
)

var (
	InternalServerError = errors.NewClass(errors.InternalServerFailure, errors.InternalServerFailure)
	BadRequestGeoChatId = errors.NewClass(errors.BadRequestIdError, errors.BadRequestIdError)
)
