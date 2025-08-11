package services

import (
	"golang-eshop-backend/internal/dto"
	"golang-eshop-backend/internal/models"
	"golang-eshop-backend/internal/api/rest/helpers/logging"

	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)


type UserService struct {}

func NewUserService() UserService {
	return UserService{}
}

func (s *UserService) FindUserByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}

func (s *UserService) SignUp(ctx *fiber.Ctx, input dto.UserSignUp) (string, error) {
	logger := logging.GetLoggerFromCtx(ctx)
	logger.Info().Msg("new user created successfully")
	
	return "my-token", nil
}

func (s *UserService) SignIn(input any) (string, error) {
	return "", nil
}

func (s *UserService) GetVerificationCode(u models.User) (int, error) {
	return 0, nil
}

func (s *UserService) VerifyUser(id uuid.UUID, code int) (string, error) {
	return "", nil
}

func (s *UserService) CreateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s *UserService) GetProfile(id uuid.UUID) (*models.User, error) {
	return &models.User{}, nil
}

func (s *UserService) UpdateProfile(id uuid.UUID, input any) (*models.User, error) {
	return &models.User{}, nil
}

func (s *UserService) BecomeSeller(id uuid.UUID, input any) (string, error) {
	return "", nil
}

func (s *UserService) GetCart(id uuid.UUID) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) UpdateCart(u models.User, input any) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) CreateOrder(u models.User) (int, error) {
	return 0, nil
}

func (s *UserService) GetOrders(u models.User) ([]interface{}, error) {
	return nil, nil
}

func (s *UserService) GetOrderById(id uint, u models.User) ([]interface{}, error) {
	return nil, nil
}
