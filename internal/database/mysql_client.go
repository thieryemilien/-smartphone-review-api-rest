package database

import (
	"database/sql"
	"fmt"
)

type MySQLClient struct {
	// Other fields like connection pool, etc.
}

func NewMySQLClient(source string) *sql.DB {

	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("Error connecting to the database: %s", err.Error())
		panic(err)
	}
	return db
}
