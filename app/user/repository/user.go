package repository

import (
	"reflect"
	"strconv"
	"umu/golang-api/app/domain"

	"gopkg.in/gorp.v2"
)

type userRepo struct {
	*gorp.DbMap
}

func NewUserRepository(db *gorp.DbMap) domain.UserRepository {
	return &userRepo{
		db,
	}
}

func (u *userRepo) Put(user *domain.User) (result *domain.User, err error) {
	err = u.Insert(user)
	if err != nil {
		return nil, err
	}
	result = user
	return result, nil
}

func (u *userRepo) Get(search domain.SearchCriteria) (result *domain.User, err error) {
	where := ""
	args := []interface{}{}

	rv := reflect.ValueOf(search)

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.IsValid() || f.IsNil() {
			continue
		}

		v := f.Elem().Interface()
		args = append(args, v)

		col := rv.Type().Field(i).Tag.Get("col")
		where += col + " = $" + strconv.Itoa(len(args)) + " OR "
	}

	if wlen := len(where); wlen > 0 {
		where = "WHERE " + where[:wlen-len(" OR ")] // prepend WHERE and drop the last AND
	}
	result = &domain.User{}
	err = u.SelectOne(result, "SELECT * FROM newsletter"+where, args...)

	return
}
