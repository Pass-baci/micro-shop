// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO user (
    id, username, password
) VALUES (?, ?, ?)
`

type CreateUserParams struct {
	ID       string
	Username string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.ID, arg.Username, arg.Password)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, create_time, update_time  FROM user
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, username, password, create_time, update_time FROM user
ORDER BY id
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.CreateTime,
			&i.UpdateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}