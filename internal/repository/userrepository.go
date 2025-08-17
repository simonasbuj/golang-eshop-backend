package repository

import (
	"golang-eshop-backend/internal/api/rest/helpers/logging"
	"golang-eshop-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx *fiber.Ctx, user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx *fiber.Ctx, user *models.User) (*models.User, error) {
	l := logging.GetLoggerFromCtx(ctx)
	l.Info().Str("email", user.Email).Msg("inserting new user into database")

	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}