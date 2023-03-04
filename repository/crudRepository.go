package repository

import (
	"context"
)

type CrudRepository[T any] interface {
	Create(ctx context.Context, t T) (*T, error)
	GetById(ctx context.Context, id int64) (*T, error)
	Close() error
}
