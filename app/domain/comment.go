package domain

import (
	"time"
	"umu/golang-api/common"
)

type Comment struct {
	Id        string            `json:"id" db:"id, primarykey"`
	Author    string            `json:"author" db:"author"`
	Content   string            `json:"content" db:"content"`
	Status    string            `json:"status" db:"status"`
	PostId    string            `json:"post_id" db:"post_id"`
	CreatedBy string            `json:"created_by" db:"created_by"`
	UpdatedBy common.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy common.NullString `json:"deleted_by" db:"deleted_by"`
	CreatedAt time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt common.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt common.NullTime   `json:"deleted_at" db:"deleted_at"`
}
