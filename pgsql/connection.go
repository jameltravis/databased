package pgsql

import (
	"database/sql"
	// Not sure if the driver is necessary but why not.
	_ "github.com/lib/pq"
)

// Conn opens a new PostgreSQL connection string.
func Conn(uname string, pass string, url string, dbname string, sslMode string) (*sql.DB, error) {
	connStr := "postgres://" + uname + ":" + pass + "@" + url + "/" + dbname + "?sslmode=" + sslMode
	return sql.Open("postgres", connStr)
}
