package usecases

import (
	"database/sql"
)

type DBRepository interface {
	BeginDB() *sql.DB
}