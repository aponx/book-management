package repository

import (
	"github.com/aponx/book-management/app/domain"
	"github.com/aponx/book-management/common"
	"github.com/aponx/book-management/driver"
	// "github.com/aponx/book-management/driver"
)

type bookRepo struct {
	db *[]domain.Book
}

func NewBookRepository(db *[]domain.Book) domain.BookRepository {
	return &bookRepo{
		db,
	}
}

func (u *bookRepo) GetAll() (result []domain.Book, err error) {
	return *u.db, common.ErrNotFoundData
}

func (u *bookRepo) GetById(bookID string) (result domain.Book, err error) {
	for _, b := range *u.db {
		if b.BookID == bookID {
			return b, nil
		}
	}
	return result, common.ErrNotFoundData
}

func (u *bookRepo) Put(book domain.Book, file string) (result domain.Book, err error) {

	data := []domain.Book{}
	for _, s := range *u.db {
		if s.BookID != book.BookID {
			data = append(data, s)
		}
	}
	data = append(data, book)

	u.db = &data

	driver.UpdateJsonFil(data, "./data/"+file)
	return book, nil
}

func (u *bookRepo) Search(search domain.SearchCriteria) (result []domain.Book, err error) {
	for _, s := range *u.db {
		if s.Title == *search.Title {
			result = append(result, s)
		} else if s.Author == *search.Author {
			result = append(result, s)
		} else if s.Publisher == *search.Publisher {
			result = append(result, s)
		}
	}

	return
}

func (u *bookRepo) Delete(bookid string, file string) error {

	data := []domain.Book{}
	for _, s := range *u.db {
		if s.BookID != bookid {
			data = append(data, s)
		}
	}

	u.db = &data

	driver.UpdateJsonFil(data, "./data/"+file)
	return nil
}
