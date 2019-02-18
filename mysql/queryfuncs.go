// Package mysql houses functions used to easily
// create query strings for use with Go's database/sql package.
package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

// QueryHelper is used as the basis for SQL functions
type QueryHelper struct {
	db *sql.DB
}

// Select func is used in place of long SQL select statement.
// takes the name of the table and a list of arguments/values
// as params. For now, returns a maximum of 100 rows. Error
// checking, defer.rows.Close() and Scan() are still needed
func (q *QueryHelper) Select(cols []string, table, filterCol, filterVal string) (*sql.Rows, error) {

	var query = `
		SELECT ` + strings.Join(cols, ", ") + `
		FROM ` + table + `
		WHERE ` + filterCol + ` = ?;`
	rows, err := q.db.Query(query, filterVal)
	return rows, err
}

// Update - updates records. For best result filter on an ID value
func (q *QueryHelper) Update(table, column, columnVal, filterCol, filterVal string) string {

	var query = `UPDATE ` + table + ` 
	SET ` + column + ` = ?
	WHERE ` + filterCol + ` = ` + filterVal + `;`

	stmt, err := q.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(columnVal)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil || rowCnt != 1 {
		log.Fatal(err)
	}
	return "Update successful"
}

// Insert inserts a single new record into the database.
func (q *QueryHelper) Insert(table string, cols []string, vals []string) (int64, error) {
	if len(vals) == 0 {
		var err = errors.New("No values given for the INSERT statement")
		return 0, err
	}
	var query = `
		INSERT INTO ` + table + ` (` + strings.Join(cols, ", ") + `)
		VALUES (?` + strings.Repeat(",?", len(vals)-1) + `);`
	stmt, err := q.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("Insert: Error while preparing statement: %s", err)
		return 0, err
	}
	v := make([]interface{}, len(vals))
	res, err := stmt.Exec(v...)
	if err != nil {
		err = fmt.Errorf("Insert: Error executing query: %v", err)
		return 0, err
	}
	getID, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("Insert: There was an issue getting the most recent ID: %v", err)
		return 0, err
	}
	rowcnt, err := res.RowsAffected()
	if err != nil || rowcnt > 1 {
		err = fmt.Errorf("The query ran with issues: error: %v rowcount: %d", err, rowcnt)
		return 0, err
	}
	return getID, nil
}

// Delete deletes a single item from the database
func (q *QueryHelper) Delete(table, filterCol, filterVal string) (int64, error) {
	var query = `
	DELETE FROM ` + table + `
	WHERE ` + filterCol + ` = ?`

	stmt, err := q.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("There was an issue preparing the query: Delete: %v", err)
		return 0, err
	}
	res, err := stmt.Exec(filterVal)
	if err != nil {
		err = fmt.Errorf("There was an issue executing the query: Delete: %v", err)
		return 0, err
	}
	getID, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("Insert: There was an issue getting the most recent ID: %v", err)
		return 0, err
	}
	rowcnt, err := res.RowsAffected()
	if err != nil || rowcnt > 1 {
		err = fmt.Errorf("The query ran with issues: error: %v rowcount: %d", err, rowcnt)
		return 0, err
	}
	return getID, nil
}
