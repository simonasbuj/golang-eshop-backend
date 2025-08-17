package services_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"golang-eshop-backend/internal/dto"
	"golang-eshop-backend/internal/models"
	"golang-eshop-backend/internal/services"
)

// mock dependencies
type mockUserRepository struct {}

func (r *mockUserRepository) CreateUser(ctx *fiber.Ctx, user *models.User) (*models.User, error) {
	return user, nil
}

func TestUserService_SignUp(t *testing.T) {
	// inti variables

    r := &mockUserRepository{}
    userService := services.NewUserService(r)

	app := fiber.New()
	mockCtx := app.AcquireCtx(&fasthttp.RequestCtx{})

    // Prepare input DTO
    input := dto.UserSignUp{
        UserSignIn: dto.UserSignIn{
            Email:    "test@example.com",
            Password: "secret",
        },
        Phone: "1234567890",
    }

	// call function
    token, err := userService.SignUp(mockCtx, input)

	// assertions
    assert.NoError(t, err)
    assert.Equal(t, "my-token", token)
}
