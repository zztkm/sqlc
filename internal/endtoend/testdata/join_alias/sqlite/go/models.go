// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	ID    int64
	Title sql.NullString
}

type Foo struct {
	ID int64
}