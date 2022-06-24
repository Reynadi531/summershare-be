package auth

import (
	"github.com/gofiber/fiber/v2"
	authService "summershare/internal/service/auth"
	"summershare/pkg/utils"
)

type authHandler struct {
	authService authService.AuthService
}

type loginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

func (a authHandler) Login(c *fiber.Ctx) error {
	login := new(loginRequest)

	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": err.Error(),
		})
	}

	errors := utils.ValidateStruct(login)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Validation failed",
			"details": errors,
		})
	}

	response := a.authService.Login(login.Username, login.Password)
	return c.Status(response.StatusCode).JSON(response)
}

type registerRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

func (a authHandler) Register(c *fiber.Ctx) error {
	register := new(registerRequest)

	if err := c.BodyParser(register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": err.Error(),
		})
	}

	errors := utils.ValidateStruct(register)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Validation failed",
			"details": errors,
		})
	}

	response := a.authService.Register(register.Username, register.Email, register.Password)
	return c.Status(response.StatusCode).JSON(response)
}

func (a authHandler) RefreshToken(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	if jwtToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": "Authorization header is required",
		})
	}

	meta, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": err.Error(),
		})
	}

	response := a.authService.RefreshToken(meta.UserID.String(), jwtToken)

	return c.Status(response.StatusCode).JSON(response)
}

func NewAuthHandler(authService authService.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}
