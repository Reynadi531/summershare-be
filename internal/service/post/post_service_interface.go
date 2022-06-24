package post

import (
	"github.com/google/uuid"
	"summershare/internal/entities/web"
	"summershare/pkg/database"
)

type PostService interface {
	Create(body string, isJoinable bool, id uuid.UUID) web.Response
	Update(body string, isJoinable bool, postId string, userId uuid.UUID) web.Response
	ViewByID(postId string) web.Response
	ViewAll(pagination *database.Pagination) web.Response
	ViewSelf(pagination *database.Pagination, userId uuid.UUID) web.Response
}
