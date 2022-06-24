package post

import (
	"summershare/internal/entities"
	"summershare/pkg/database"
)

type PostRepository interface {
	Create(post entities.Post) error
	Update(post entities.Post) error
	GetByID(id string) (entities.Post, error)
	GetAll(pagination database.Pagination) (*database.Pagination, []*entities.Post, error)
	GetByOwnerId(pagination database.Pagination, ownerId string) (*database.Pagination, []*entities.Post, error)
	CountByUserID(userID string) (int64, error)
}
