package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conn establishes a connection to a MySQL database.
// Note that a `defer()` call is still needed after `MySQLConn()`
// is called.
func Conn(uname string, psswd string, prtcl string, ip string, schema string) (*sql.DB, error) {
	connStr := uname + ":" + psswd + "@" + prtcl + "(" + ip + ")/" + schema
	return sql.Open("mysql", connStr)
}

// #TODO PGConn func

// #TODO MSSQLConn func
