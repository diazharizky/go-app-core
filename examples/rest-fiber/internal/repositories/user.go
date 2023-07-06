package repositories

import (
	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{db}
}

func (repo userRepository) List() (users []models.User, err error) {
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

func (repo userRepository) Create(newUser *models.User) error {
	return repo.db.Create(newUser).Error
}
