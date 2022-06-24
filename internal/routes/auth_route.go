package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	authHandler "summershare/internal/handler/auth"
	userRepo "summershare/internal/repository/user"
	authService "summershare/internal/service/auth"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	authService := authService.NewAuthService(userRepo)
	authHandler := authHandler.NewAuthHandler(authService)

	authRouteGroup := app.Group("/api/v1/auth")
	authRouteGroup.Post("/login", authHandler.Login)
	authRouteGroup.Post("/register", authHandler.Register)
	authRouteGroup.Post("/refresh", authHandler.RefreshToken)
}
