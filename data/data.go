// Package data provides structure and interaction handling
// for interfacing with a database.
package data

import (
	"github.com/jmoiron/sqlx"
)

// Data is a structure to help work with database connections
// and abstract a useful interface for consumers to utilise.
type Data struct {
	Errors []error
	c      *sqlx.DB
}

// Open will setup the sqlite file and db handle.
func Open() (*Data, error) {
	db, err := sqlx.Open("sqlite3", "./.go-vue-skeleton.db")
	if err != nil {
		return nil, err
	}

	d := &Data{c: db}
	err = d.migrate()
	return d, err
}
