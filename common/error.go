package common

import (
	"errors"
	"net/http"

	phttp "github.com/aponx/book-management/common/http"
)

var ErrInvalidRequest = errors.New("invalid request")
var ErrUnprocessableEntity = errors.New("unprocessable data")
var ErrNotFoundData = errors.New("data not found")
var ErrBookAlreadyExist = errors.New("book already exist")
var ErrUserNotFound = errors.New("user not found")
var ErrPasswordNotMatch = errors.New("password doesn't match")

func InjectErrors(handlerCtx *phttp.HttpHandlerContext) {
	handlerCtx.AddError(ErrNotFoundData, setErrResp(ErrNotFoundData.Error(), http.StatusNotFound))
	handlerCtx.AddError(ErrUnprocessableEntity, setErrResp(ErrUnprocessableEntity.Error(), http.StatusUnprocessableEntity))
	handlerCtx.AddError(ErrInvalidRequest, setErrResp(ErrInvalidRequest.Error(), http.StatusBadRequest))
	handlerCtx.AddError(ErrBookAlreadyExist, setErrResp(ErrBookAlreadyExist.Error(), http.StatusBadRequest))
	handlerCtx.AddError(ErrUserNotFound, setErrResp(ErrUserNotFound.Error(), http.StatusNotFound))
	handlerCtx.AddError(ErrPasswordNotMatch, setErrResp(ErrPasswordNotMatch.Error(), http.StatusConflict))

}

func setErrResp(message string, statusCode int) *phttp.ErrorResponse {
	return &phttp.ErrorResponse{
		Response: phttp.Response{
			ResponseDesc: message,
		},
		HttpStatus: statusCode,
	}
}
