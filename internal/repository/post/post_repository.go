package post

import (
	"gorm.io/gorm"
	"summershare/internal/entities"
	"summershare/pkg/database"
)

type postRepository struct {
	DB *gorm.DB
}

func (p postRepository) GetByOwnerId(pagination database.Pagination, ownerId string) (*database.Pagination, []*entities.Post, error) {
	var posts []*entities.Post
	p.DB.Scopes(database.Paginate(posts, &pagination, p.DB)).Preload("Owner").Find(&posts)
	return &pagination, posts, nil
}

func (p postRepository) Create(post entities.Post) error {
	return p.DB.Create(&post).Error
}

func (p postRepository) Update(post entities.Post) error {
	return p.DB.Model(&post).Updates(post).Error
}

func (p postRepository) GetByID(id string) (entities.Post, error) {
	var post entities.Post
	err := p.DB.Where("id = ?", id).Preload("Owner").First(&post).Error
	return post, err
}

func (p postRepository) GetAll(pagination database.Pagination) (*database.Pagination, []*entities.Post, error) {
	var posts []*entities.Post
	p.DB.Scopes(database.Paginate(posts, &pagination, p.DB)).Preload("Owner").Find(&posts)
	return &pagination, posts, nil
}

func (p postRepository) CountByUserID(userID string) (int64, error) {
	var count int64
	err := p.DB.Model(&entities.Post{}).Where("owner_id = ?", userID).Count(&count).Error
	return count, err
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		DB: db,
	}
}
