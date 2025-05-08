package http

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"

	"github.com/aponx/book-management/app/domain"
	"github.com/aponx/book-management/common"
	phttp "github.com/aponx/book-management/common/http"
)

type bookHandler struct {
	*common.Config
	domain.BookUsecase
}

func NewBookHandler(conf *common.Config, bookUsecase domain.BookUsecase) http.Handler {
	handler := &bookHandler{
		conf,
		bookUsecase,
	}

	handlerCtx := phttp.NewContextHandler(conf.Log.IsDebug())

	common.InjectErrors(&handlerCtx)

	r := chi.NewRouter()

	phandler := phttp.NewHttpHandler(handlerCtx)

	r.Method(http.MethodGet, "/", phandler(handler.GetAll))
	r.Method(http.MethodGet, "/{id}", phandler(handler.GetById))
	r.Method(http.MethodPost, "/search", phandler(handler.Search))
	r.Method(http.MethodPost, "/", phandler(handler.Create))
	r.Method(http.MethodPut, "/{id}", phandler(handler.Update))
	r.Method(http.MethodDelete, "/{id}", phandler(handler.Delete))

	return r

}

func (h *bookHandler) GetAll(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	// get all book to usecase
	books, err := h.BookUsecase.GetAll()
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusOK
	response.Data = books
	response.Msg = "successfully found data"
	return
}

func (h *bookHandler) GetById(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	bookid := chi.URLParam(r, "id")
	// get by id book to usecase
	books, err := h.BookUsecase.GetById(bookid)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusOK
	response.Data = books
	response.Msg = "successfully found data"
	return
}

func (h *bookHandler) Search(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	// Read from Request body
	var searchReq domain.SearchCriteria
	err := json.NewDecoder(r.Body).Decode(&searchReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}
	// Validate request
	_, err = govalidator.ValidateStruct(searchReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}

	// get book to usecase
	books, err := h.BookUsecase.GetBookByCriteria(searchReq)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusOK
	response.Data = books
	response.Msg = "successfully found data"
	return
}

func (h *bookHandler) Create(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	// Read from Request body
	var createReq domain.BookCreateRequest
	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}
	// Validate request
	_, err = govalidator.ValidateStruct(createReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}

	// Create book
	book, err := h.BookUsecase.Create(createReq, h.Config.JSON.Data)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusCreated
	response.Data = book
	response.Msg = "successfully created data"
	return
}

func (h *bookHandler) Update(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	bookid := chi.URLParam(r, "id")
	// Read from Request body
	var updateReq domain.BookUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}
	// Validate request
	_, err = govalidator.ValidateStruct(updateReq)
	if err != nil {
		response.Error = common.ErrInvalidRequest
		return
	}

	// Update book
	book, err := h.BookUsecase.Update(bookid, updateReq, h.Config.JSON.Data)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusOK
	response.Data = book
	response.Msg = "successfully updated data"
	return
}

func (h *bookHandler) Delete(w http.ResponseWriter, r *http.Request) (response phttp.HttpHandleResult) {
	bookid := chi.URLParam(r, "id")

	// Delete book
	err := h.BookUsecase.Delete(bookid, h.Config.JSON.Data)
	if err != nil {
		response.Error = err
		return
	}
	response.StatusCode = http.StatusOK
	response.Msg = "successfully deleted data"
	return
}
