package handlers

import (
	"golang-eshop-backend/internal/api/rest"
	"golang-eshop-backend/internal/dto"
	"golang-eshop-backend/internal/helpers/logging"
	"golang-eshop-backend/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	s 		services.UserService
	logger 	*zerolog.Logger
}

func newUserHandler(service services.UserService, logger *zerolog.Logger) *UserHandler {
	l := logger.With().Str("class", "UserHandler").Logger()
	
	return &UserHandler{
		s: service,
		logger: &l,
	}
}

func SetupUserRoutes(rh *rest.RestHandler, logger *zerolog.Logger) {
	app := rh.App

	userService := services.NewUserService(logger)
	h := newUserHandler(userService, logger)

	// public endpoints
	app.Post("/auth/signup", h.SignUp)
	app.Post("/auth/signin", h.SignIn)

	// private endpoints
	app.Get("/auth/verify", h.Verify)
	app.Post("/auth/verify", h.GetVerificationCode)

	app.Get("/user/profile", h.GetProfile)
	app.Post("/user/profile", h.UpdateProfile)

	app.Get("/cart", h.GetCart)
	app.Post("/cart", h.AddToCart)
	app.Get("/orders", h.GetOrders)
	app.Get("/orders/:id", h.GetOrder)

	app.Post("/become-seller", h.SignIn)
}


func (h *UserHandler) SignUp(ctx *fiber.Ctx) error {
	logging.LogInfo(h.logger, ctx, "singup handler triggered")
	logger := logging.GetLoggerFromCtx(ctx)
	logger.Info().Msg("does this finally work?")
	
	// try to create user from passed json
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		errorMsg := "invalid user input provided"
		logging.LogError(h.logger, ctx, err, errorMsg)
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": errorMsg,
		})
	}

	token, err := h.s.SignUp(ctx, user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "error during signup",
		})
	}

	logging.LogInfo(h.logger, ctx, "singup handler finished succesfully")
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "new user signed up successfully",
		"token": token,
	})
}

func (h *UserHandler) SignIn(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "sing in",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verify",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getting verification code",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getting profile",
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "creating profile",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getting cart",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "adding to cart",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getting orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "getting a particular order",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "creating order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
	})
}
