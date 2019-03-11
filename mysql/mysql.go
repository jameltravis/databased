package mysql

import "database/sql"

// IQuery is used for DB access. Each method corresponds to the associated
// sql statement.
type IQuery interface {
	Select() (*sql.Rows, error)
	Update() string
	Insert() (int64, error)
	Delete() (int64, error)
}

// IModel is an interface for models. Used so struct models are able to call
// the methods below. To use this interface add Model as an embedded
// field in a struct.
type IModel interface {
	IQuery
	ListItems() (*sql.Rows, error)
	LreateItem() (int64, error)
	UpdateItem() string
	DeleteItem() (int64, error)
}
