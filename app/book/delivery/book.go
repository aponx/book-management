package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/aponx/book-management/app/domain"
	"github.com/aponx/book-management/common"
	phttp "github.com/aponx/book-management/common/http"
)

type bookHandler struct {
	*common.Config
	domain.BookUsecase
}

func NewUserHandler(conf *common.Config, bookUsecase domain.BookUsecase) http.Handler {
	handlerCtx := phttp.NewContextHandler(conf.Log.IsDebug())

	common.InjectErrors(&handlerCtx)

	r := chi.NewRouter()

	phandler := phttp.NewHttpHandler(handlerCtx)

	// r.Method(http.MethodPost, "/register", phandler(handler.Register))

	return r

}

// func (h *bookHandler) Register(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
// 	// Read from Request body
// 	var userReq domain.UserRegisterRequest
// 	err := json.NewDecoder(r.Body).Decode(&userReq)
// 	if err != nil {
// 		response.Error = common.ErrInvalidRequest
// 		return
// 	}
// 	// Validate request
// 	_, err = govalidator.ValidateStruct(userReq)
// 	if err != nil {
// 		response.Error = common.ErrInvalidRequest
// 		return
// 	}

// 	// Register user
// 	user, err := h.BookUsecase.Register(&userReq)
// 	if err != nil {
// 		response.Error = err
// 		return
// 	}
// 	response.StatusCode = http.StatusCreated
// 	response.Data = user
// 	return
// }
