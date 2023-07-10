package app

import (
	"context"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
)

type IUserRepository interface {
	Create(ctx context.Context, newUser *models.User) error
	List(ctx context.Context) ([]models.User, error)
}
