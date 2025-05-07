package usecase

import (
	"github.com/aponx/book-management/app/domain"
)

type bookUsecase struct {
}

func NewBookUsecase() domain.BookUsecase {
	return &bookUsecase{}
}
