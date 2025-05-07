package domain

import "net/http"

type Repository struct {
	BookRepo BookRepository
}

type Usecase struct {
	BookUsecase BookUsecase
}

type Delivery struct {
	BookDelivery http.Handler
}
