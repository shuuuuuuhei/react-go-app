package infrastructure

import (
	"github.com/react-go-app/backend/interfaces/database"
	"database/sql"
)

type SqlHandler struct {
	db *sql.DB
}

type Result struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
}

func NewSqlHandler() database.SqlHandler {
	dsn := "sample code"
	db, err := sql.Open("postgres", dsn)
	defer db.Close()

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Execute(paramater string, args ...interface{}) (database.SqlResult, error) {
	res := SqlResult{}
	result, err := handler.db.Exec(paramater, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(paramater string, args ...interface{}) (database.SqlRow, error) {
	rows, err := handler.db.Query(paramater, args)
	
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}