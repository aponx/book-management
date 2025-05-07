package repository

import (
	"github.com/aponx/book-management/app/domain"
)

type bookRepo struct {
	*domain.Book
}

func NewBookRepository(db *[]domain.Book) domain.BookRepository {
	return &bookRepo{
		db,
	}
}

func (u *bookRepo) Put(book *domain.Book) (result *domain.Book, err error) {
	err = u.Insert(book)
	if err != nil {
		return nil, err
	}
	result = book
	return result, nil
}

func (u *bookRepo) Get(search domain.SearchCriteria) (result *domain.Book, err error) {

	return
}
