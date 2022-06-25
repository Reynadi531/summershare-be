package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	postHandler "summershare/internal/handler/post"
	postRepo "summershare/internal/repository/post"
	postService "summershare/internal/service/post"
	"summershare/pkg/middleware"
)

func RegisterPostRoute(app *fiber.App, db *gorm.DB) {
	postRepository := postRepo.NewPostRepository(db)
	postService := postService.NewPostService(postRepository)
	postHandler := postHandler.NewPostHandler(postService)

	postRouteGroup := app.Group("/api/v1/post")
	postRouteGroup.Get("/me", postHandler.ViewSelf)
	postRouteGroup.Post("/", middleware.JWTProtected(), postHandler.Create)
	postRouteGroup.Post("/:id", middleware.JWTProtected(), postHandler.Update)
	postRouteGroup.Get("/:id", postHandler.ViewByID)
	postRouteGroup.Get("/", postHandler.ViewAll)
}
