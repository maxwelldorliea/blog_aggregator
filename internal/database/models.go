// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	ApiKey    string
}