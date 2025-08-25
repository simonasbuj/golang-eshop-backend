package repository

import (
	"errors"
	"golang-eshop-backend/internal/api/rest/helpers/logging"
	"golang-eshop-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(ctx *fiber.Ctx, user *models.User) (*models.User, error)
	FindUserByEmail(ctx *fiber.Ctx, email string) (*models.User, error)
	FindUserById(ctx *fiber.Ctx, id uuid.UUID) (*models.User, error)
	UpdateUser(id uuid.UUID, u *models.User) (*models.User, error)
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

func (r *userRepository) FindUserByEmail(ctx *fiber.Ctx, email string) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		return &models.User{}, errors.New("user does not exist")
	}

	return &user, nil
}

func (r *userRepository) FindUserById(ctx *fiber.Ctx, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return &models.User{}, errors.New("user does not exist")
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(id uuid.UUID, u *models.User) (*models.User, error) {
	var user models.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		return &models.User{}, errors.New("failed to update user")
	}

	return &user, nil
}
