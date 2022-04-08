package driver

import (
	"database/sql"
	"fmt"

	"umu/golang-api/common"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

// NewPostgreDatabase return gorp dbmap object with postgre options param
func NewPostgreDatabase(option common.DB) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", option.Host, option.Port, option.Username, option.Name, option.Password))
	if err != nil {
		return nil, err
	}

	gorp := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return gorp, nil
}
