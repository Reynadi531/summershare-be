package user

import (
	"gorm.io/gorm"
	"summershare/internal/entities"
	"summershare/pkg/database"
)

type userRepository struct {
	DB *gorm.DB
}

func (u *userRepository) Update(user entities.User) error {
	return u.DB.Model(&user).Updates(user).Error
}

func (u userRepository) Create(user entities.User) error {
	return u.DB.Create(&user).Error
}

func (u userRepository) GetByID(id string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u userRepository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u userRepository) GetByUsername(username string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u userRepository) GetAll(pagination database.Pagination) (*database.Pagination, []*entities.User, error) {
	var users []*entities.User
	err := u.DB.Scopes(database.Paginate(users, &pagination, u.DB)).Find(&users).Error
	return &pagination, users, err
}

func (u userRepository) CountByUsername(username string) (int64, error) {
	var count int64
	err := u.DB.Model(&entities.User{}).Where("username = ?", username).Count(&count).Error
	return count, err
}

func (u userRepository) CountByEmail(email string) (int64, error) {
	var count int64
	err := u.DB.Model(&entities.User{}).Where("email = ?", email).Count(&count).Error
	return count, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}
