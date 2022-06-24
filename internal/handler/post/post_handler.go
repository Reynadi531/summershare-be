package post

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	postService "summershare/internal/service/post"
	"summershare/pkg/database"
	"summershare/pkg/utils"
)

type postHandler struct {
	postService postService.PostService
}

// GET /api/v1/post/me?page=1&limit=10&sort=created_at&order=desc
func (p postHandler) ViewSelf(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")
	order := c.Query("order")

	pagination := database.Pagination{
		Page:  page,
		Limit: limit,
		Sort:  fmt.Sprintf("%s %s", sort, order),
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

	response := p.postService.ViewSelf(&pagination, meta.UserID)
	return c.Status(response.StatusCode).JSON(response)
}

type createRequest struct {
	Body       string `json:"body" validate:"required,min=30,max=2000"`
	IsJoinable bool   `json:"is_joinable"`
}

// POST /api/v1/post
func (p postHandler) Create(c *fiber.Ctx) error {
	createBody := new(createRequest)

	if err := c.BodyParser(createBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": err.Error(),
		})
	}

	errors := utils.ValidateStruct(createBody)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Validation failed",
			"details": errors,
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

	response := p.postService.Create(createBody.Body, createBody.IsJoinable, meta.UserID)

	return c.Status(response.StatusCode).JSON(response)
}

// POST /api/v1/post/{post_id}
func (p postHandler) Update(c *fiber.Ctx) error {
	postId := c.Params("id")
	if postId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": "post_id is required",
		})
	}

	updateBody := new(createRequest)
	if err := c.BodyParser(updateBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": err.Error(),
		})
	}

	errors := utils.ValidateStruct(updateBody)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Validation failed",
			"details": errors,
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

	response := p.postService.Update(updateBody.Body, updateBody.IsJoinable, postId, meta.UserID)

	return c.Status(response.StatusCode).JSON(response)
}

// GET /api/v1/post/{post_id}
func (p postHandler) ViewByID(c *fiber.Ctx) error {
	postId := c.Params("id")
	if postId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"message": "Invalid request body",
			"details": "post_id is required",
		})
	}

	response := p.postService.ViewByID(postId)
	return c.Status(response.StatusCode).JSON(response)
}

// GET /api/v1/post?page=1&limit=10&sort=created_at&order=desc
func (p postHandler) ViewAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")
	order := c.Query("order")

	pagination := database.Pagination{
		Page:  page,
		Limit: limit,
		Sort:  fmt.Sprintf("%s %s", sort, order),
	}

	response := p.postService.ViewAll(&pagination)
	return c.Status(response.StatusCode).JSON(response)
}

func NewPostHandler(postService postService.PostService) PostHandler {
	return &postHandler{
		postService: postService,
	}
}
