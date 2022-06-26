package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"summershare/internal/entities"
	"summershare/internal/entities/web"
	postRepo "summershare/internal/repository/post"
	"summershare/pkg/database"
	"summershare/pkg/utils"
	"time"
)

type postService struct {
	postRepository postRepo.PostRepository
}

type SendablePost struct {
	ID         int    `json:"id"`
	Body       string `json:"body"`
	IsJoinable bool   `json:"is_joinable"`
	User       struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p postService) ViewSelf(pagination *database.Pagination, userId uuid.UUID) web.Response {
	pageInfo, posts, err := p.postRepository.GetByOwnerId(*pagination, userId.String())
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to get the posts",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	var sendablePosts []SendablePost
	for _, post := range posts {
		sendablePosts = append(sendablePosts, SendablePost{
			ID:         post.ID,
			Body:       post.Body,
			IsJoinable: post.IsJoinable,
			User: struct {
				ID       string `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
			}{
				ID:       post.Owner.ID.String(),
				Username: post.Owner.Username,
				Email:    post.Owner.Email,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Posts found successfully",
		Data:       map[string]interface{}{"pageInfo": pageInfo, "posts": sendablePosts},
		Error:      nil,
	}
}

func (p postService) ViewAll(pagination *database.Pagination) web.Response {
	pageInfo, posts, err := p.postRepository.GetAll(*pagination)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to get the posts",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	var sendablePosts []SendablePost
	for _, post := range posts {
		sendablePosts = append(sendablePosts, SendablePost{
			ID:         post.ID,
			Body:       post.Body,
			IsJoinable: post.IsJoinable,
			User: struct {
				ID       string `json:"id"`
				Username string `json:"username"`
				Email    string `json:"email"`
			}{
				ID:       post.Owner.ID.String(),
				Username: post.Owner.Username,
				Email:    post.Owner.Email,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Posts found successfully",
		Data:       map[string]interface{}{"pageInfo": pageInfo, "posts": sendablePosts},
		Error:      nil,
	}
}

func (p postService) ViewByID(postId string) web.Response {
	post, err := p.postRepository.GetByID(postId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to get the post",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	if err == gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "Post not found",
			Data:       nil,
			Error:      fiber.ErrNotFound,
		}
	}

	var sendablePost SendablePost
	sendablePost = SendablePost{
		ID:         post.ID,
		Body:       post.Body,
		IsJoinable: post.IsJoinable,
		User: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       post.Owner.ID.String(),
			Username: post.Owner.Username,
			Email:    post.Owner.Email,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Post found successfully",
		Data:       sendablePost,
		Error:      nil,
	}
}

func (p postService) Update(body string, isJoinable bool, postId string, userId uuid.UUID) web.Response {
	post, err := p.postRepository.GetByID(postId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to get the post",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	if err == gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "Post not found",
			Data:       nil,
			Error:      fiber.ErrNotFound,
		}
	}

	if post.OwnerID != userId {
		return web.Response{
			StatusCode: fiber.StatusForbidden,
			Message:    "You are not allowed to update this post",
			Data:       nil,
			Error:      fiber.ErrForbidden,
		}
	}

	post.Body = body
	post.IsJoinable = isJoinable
	post.UpdatedAt = time.Now()

	if err := p.postRepository.Update(post); err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to update the post",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	var sendablePost SendablePost
	sendablePost = SendablePost{
		ID:         post.ID,
		Body:       post.Body,
		IsJoinable: post.IsJoinable,
		User: struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       post.Owner.ID.String(),
			Username: post.Owner.Username,
			Email:    post.Owner.Email,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Post updated successfully",
		Data:       sendablePost,
		Error:      nil,
	}
}

func (p postService) Create(body string, isJoinable bool, id uuid.UUID) web.Response {
	post := entities.Post{
		ID:         utils.GeneratePostId(),
		Body:       body,
		IsJoinable: isJoinable,
		OwnerID:    id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := p.postRepository.Create(post); err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Failed to save the post",
			Data:       nil,
			Error:      fiber.ErrInternalServerError,
		}
	}

	return web.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "Post created successfully",
		Data:       post,
		Error:      nil,
	}
}

func NewPostService(postRepository postRepo.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}
