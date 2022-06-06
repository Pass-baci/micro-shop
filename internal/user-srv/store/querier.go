// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package store

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (User, error)
	ListUser(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)