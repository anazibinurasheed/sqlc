// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID       int64
	Name     string
	ParentID sql.NullInt64
}