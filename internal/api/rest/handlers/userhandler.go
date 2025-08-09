package handlers

import (
	"golang-eshop-backend/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// TODO pass UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := UserHandler{}

	// public endpoints
	app.Post("/auth/signup", h.SignUp)
	app.Post("/auth/signin", h.SignIn)

	// private endpoints
	app.Get("/auth/verify", h.Verify)
	app.Post("/auth/verify", h.GetVerificationCode)
	app.Get("/auth/profile", h.GetProfile)
	app.Post("/auth/profile", h.CreateProfile)

	app.Get("/cart", h.GetCart)
	app.Post("/cart", h.AddToCart)
	app.Get("/orders", h.GetOrders)
	app.Get("/orders/:id", h.GetOrder)

	app.Post("/become-seller", h.SignIn)
}


func (h *UserHandler) SignUp(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "sing up",
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

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
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
