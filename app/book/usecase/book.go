package usecase

import (
	"github.com/aponx/book-management/app/domain"
	"github.com/aponx/book-management/common"
)

type bookUsecase struct {
	bookRepo domain.BookRepository
}

func NewBookUsecase(bookRepo domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (u *bookUsecase) GetAll() (books []domain.Book, err error) {
	books, err = u.bookRepo.GetAll()
	return
}

func (u *bookUsecase) GetById(bookID string) (books domain.Book, err error) {
	books, err = u.bookRepo.GetById(bookID)
	return
}

func (u *bookUsecase) GetBookByCriteria(search domain.SearchCriteria) (books []domain.Book, err error) {
	books, err = u.bookRepo.Search(search)
	return
}

func (u *bookUsecase) Create(req domain.BookCreateRequest, file string) (books domain.Book, err error) {
	_, err = u.bookRepo.GetById(req.BookID)
	if err == nil {
		return books, common.ErrBookAlreadyExist
	}

	data := domain.Book{
		BookID:    req.BookID,
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
		Qty:       req.Qty,
		Out:       0,
	}

	books, err = u.bookRepo.Put(data, file)
	if err != nil {
		return books, common.ErrUnprocessableEntity
	}
	return
}

func (u *bookUsecase) Update(bookID string, req domain.BookUpdateRequest, file string) (books domain.Book, err error) {
	_, err = u.bookRepo.GetById(bookID)
	if err != nil {
		return books, common.ErrNotFoundData
	}

	if req.Out > req.Qty {
		return books, common.ErrInvalidRequest
	}

	data := domain.Book{
		BookID:    bookID,
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
		Year:      req.Year,
		Qty:       req.Qty,
		Out:       req.Out,
	}

	books, err = u.bookRepo.Put(data, file)
	if err != nil {
		return books, common.ErrUnprocessableEntity
	}
	return
}

func (u *bookUsecase) Delete(bookID string, file string) error {
	_, err := u.bookRepo.GetById(bookID)
	if err != nil {
		return common.ErrNotFoundData
	}

	err = u.bookRepo.Delete(bookID, file)
	if err != nil {
		return common.ErrUnprocessableEntity
	}
	return nil
}
