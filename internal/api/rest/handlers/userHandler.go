package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app2/internal/api/rest"
	"go-ecommerce-app2/internal/dto"
	"go-ecommerce-app2/internal/repository"
	"go-ecommerce-app2/internal/service"
	"log"
)

type UserHandler struct {
	// svc
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}
	// create handler having related routes in it
	handler := UserHandler{svc}

	pubRoutes := app.Group("/users")
	//public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	//private endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)
	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)
	pvtRoutes.Post("/become-seller", handler.BecomeSeller)
}

func (uh UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
		})
	}
	token, err := uh.svc.SignUp(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})
}

func (uh UserHandler) Login(ctx *fiber.Ctx) error {
	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := uh.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "please provide correct email and password",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}

func (uh UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := uh.svc.Auth.GetCurrentUser(ctx)
	code, err := uh.svc.GetVerificationCode(user)
	if err != nil {
		if err.Error() == "user already verified" {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": "user already verified",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "verification code generated",
		"data":    code,
	})
}

func (uh UserHandler) Verify(ctx *fiber.Ctx) error {

	user := uh.svc.Auth.GetCurrentUser(ctx)

	// validate incoming request
	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid input",
		})
	}
	err := uh.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"err": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}

func (uh UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := uh.svc.Auth.GetCurrentUser(ctx)
	log.Println(user)
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "profile",
		"user":    user,
	})
}

func (uh UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "create profile",
	})
}

func (uh UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (uh UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
	})
}

func (uh UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (uh UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}

func (uh UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
	})
}
