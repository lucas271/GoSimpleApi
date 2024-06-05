// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one

select id, created_at, updated_at, name from users
`

func (q *Queries) CreateUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}