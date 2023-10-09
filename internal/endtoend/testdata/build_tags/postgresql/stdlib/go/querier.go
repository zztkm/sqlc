//go:build some_tag

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package authors

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Author, error)
	ListAuthors(ctx context.Context) ([]Author, error)
}

var _ Querier = (*Queries)(nil)