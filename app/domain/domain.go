package domain

import "net/http"

type Repository struct {
	UserRepo UserRepository
}

type Usecase struct {
	UserUsecase UserUsecase
}

type Delivery struct {
	UserDelivery http.Handler
}
