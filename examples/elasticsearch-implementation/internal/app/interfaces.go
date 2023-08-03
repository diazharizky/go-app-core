package app

import (
	"context"

	"github.com/diazharizky/go-app-core/examples/elasticsearch-implementation/internal/models"
)

type IUserRepository interface {
	List(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, newUser models.User) error
}
