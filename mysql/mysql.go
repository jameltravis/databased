package mysql

import "database/sql"

// IModel is an interface for models. Used so struct models are able to call
// the methods below. To use this interface add Model as an embedded
// field in a struct.
type IModel interface {
	listItems() (*sql.Rows, error)
	createItem() (int64, error)
	updateItem() string
	deleteItem() (int64, error)
}
