package domain

import (
	"time"
	"umu/golang-api/common"
)

type SearchCriteria struct {
	Username *string `col:"username"`
	Email    *string `col:"email"`
	Phone    *string `col:"phone"`
}

type User struct {
	Id        string            `json:"id" db:"id, primarykey"`
	Name      string            `json:"name" db:"name"`
	Username  string            `json:"username" db:"username"`
	Email     string            `json:"email" db:"email"`
	Phone     string            `json:"phone" db:"phone"`
	Password  string            `json:"password" db:"password"`
	CreatedBy string            `json:"created_by" db:"created_by"`
	UpdatedBy common.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy common.NullString `json:"deleted_by" db:"deleted_by"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt common.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt common.NullTime   `json:"deleted_at" db:"deleted_at"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	Put(user *User) (result *User, err error)
	Get(search SearchCriteria) (result *User, err error)
}

type UserUsecase interface {
	Register(user *UserRegisterRequest) (result *User, err error)
}

func (p *UserRegisterRequest) FillModel() (user *User) {
	timeNow := time.Now()

	user = &User{
		Name:      p.Name,
		Username:  p.Username,
		Email:     p.Email,
		Phone:     p.Phone,
		Password:  common.HashAndSalt([]byte(p.Password)),
		CreatedAt: timeNow,
		CreatedBy: "Root",
	}
	return
}
