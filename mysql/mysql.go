package mysql

import "database/sql"

//  List, Get, Create, Update, and Delete

// APIActions are the standard actions used in resource oriented
// design. This package assumes that you are adhering to this
// design strategy.
type APIActions interface {
	listItems() (*sql.Rows, error)
	createItem() (int64, error)
	updateItem() string
	deleteItem() (int64, error)
}

// Model forms the basis of normal models used by the user
// package. Models should be used as an embeded type in
// Models outside of this package.
type Model struct {
	db      *sql.DB
	success bool
}
