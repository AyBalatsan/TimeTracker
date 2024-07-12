package repository

import "github.com/go-sqlx/sqlx"

type Repository struct {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
