package repositories

import (
	"context"

	"github.com/diazharizky/go-app-core/examples/rest-fiber/internal/models"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type userRepository struct {
	traceName string
	db        *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		traceName: "repositories.user",
		db:        db,
	}
}

func (repo userRepository) List(ctx context.Context) (users []models.User, err error) {
	ctx, span := otel.Tracer(repo.traceName).Start(ctx, "list")

	// Set span's attributes here

	defer span.End()

	if err := repo.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return
}

func (repo userRepository) Create(ctx context.Context, newUser *models.User) error {
	ctx, span := otel.Tracer(repo.traceName).Start(ctx, "create")

	// Set span's attributes here

	defer span.End()

	return repo.db.WithContext(ctx).Create(newUser).Error
}
