package app

import "github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"

type IUserRepository interface {
	Create(newUser *models.User) error
	List() ([]models.User, error)
}
