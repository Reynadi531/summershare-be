package user

import (
	"summershare/internal/entities"
	"summershare/pkg/database"
)

type UserRepository interface {
	Update(user entities.User) error
	Create(user entities.User) error
	GetByID(id string) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	GetByUsername(username string) (entities.User, error)
	GetAll(pagination database.Pagination) (*database.Pagination, []*entities.User, error)
	CountByUsername(username string) (int64, error)
	CountByEmail(email string) (int64, error)
}
