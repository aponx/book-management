package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"

	"umu/golang-api/app/domain"
	"umu/golang-api/common"
	phttp "umu/golang-api/common/http"
)

type userHandler struct {
	*common.Config
	domain.UserUsecase
}

func NewUserHandler(conf *common.Config, userUsecase domain.UserUsecase) http.Handler {
	handler := &userHandler{
		conf,
		userUsecase,
	}
	handlerCtx := phttp.NewContextHandler(conf.Log.IsDebug())

	common.InjectErrors(&handlerCtx)

	r := chi.NewRouter()

	phandler := phttp.NewHttpHandler(handlerCtx)

	r.Method(http.MethodPost, "/register", phandler(handler.Register))
	r.Method(http.MethodGet, "/login", phandler(handler.Login))

	return r

}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	// Read from Request body
	var userReq domain.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}
	// Validate request
	_, err = govalidator.ValidateStruct(userReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}

	// Register user
	user, err := h.UserUsecase.Register(&userReq)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusCreated
	response.Data = user
	return
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	// Read from Request body
	var userReq domain.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}
	// Validate request
	_, err = govalidator.ValidateStruct(userReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}

	return
}
