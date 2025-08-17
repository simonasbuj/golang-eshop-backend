package services

import (
	"golang-eshop-backend/internal/api/rest/helpers/logging"
	"golang-eshop-backend/internal/dto"
	"golang-eshop-backend/internal/models"
	"golang-eshop-backend/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) FindUserByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}

func (s *UserService) SignUp(ctx *fiber.Ctx, input dto.UserSignUp) (string, error) {
	logger := logging.GetLoggerFromCtx(ctx).With().Str("email", input.Email).Logger()
	logger.Info().Msgf("starting user signup process")

	_, err := s.repo.CreateUser(ctx, &models.User{
		Email: input.Email,
		Password: input.Password,
		Phone: input.Phone,
	})
	if err != nil {
		return "", err
	}

	// generate token
	token := "my-token" 
	logger.Info().Msgf("new user created successfully")
	return token, nil
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
