package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Registration interface {
	RegisterUser(token string, uid int) error
	UpdateUser(token string, uid int) error
}

type Repository struct {
	Registration
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Registration: NewRegistrationMysql(db),
	}
}
