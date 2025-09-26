package helper

import (
	"errors"
	"golang-eshop-backend/internal/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func (a *Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (a *Auth) GenerateToken(id uuid.UUID, email string, role string) (string, error) {
	if id == uuid.Nil || email == "" || role == "" {
		return "", errors.New("invalid id, email or role provided")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    id,
		"email":      email,
		"role":       role,
		"expires_at": time.Now().Add(time.Hour * 24 * 30),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (a *Auth) VerifyPassword(plainPassword string, hashedPassword string) error {
	if len(plainPassword) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}

func (a *Auth) VerifyToken(token string) (models.User, error) {
	// token example: "Bearer t1234123412341234"
	tokenArr := strings.Split(token, " ")
	if len(tokenArr) != 2 {
		return models.User{}, errors.New("invalid token")
	}

	if tokenArr[0] != "Bearer" {
		return models.User{}, errors.New("invalid token")
	}

	return models.User{}, nil
}

func (a *Auth) Authorize(ctx *fiber.Ctx) error {
	return nil
}

func (a *Auth) CurrentUser(ctx *fiber.Ctx) (models.User, error) {
	return models.User{}, nil
}
