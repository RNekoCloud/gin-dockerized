// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package sqlc

import (
	"database/sql"
)

type Movie struct {
	ID        int32        `json:"id"`
	Title     string       `json:"title"`
	Genre     string       `json:"genre"`
	CreatedAt sql.NullTime `json:"created_at"`
}
