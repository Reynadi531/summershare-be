package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Email        string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Username     string    `gorm:"type:varchar(255);unique_index;not_null" json:"username"`
	Password     string    `gorm:"type:varchar(255);not_null" json:"password"`
	RefreshToken string    `gorm:"type:varchar(255)" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:current_timestamp;not_null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:current_timestamp;not_null" json:"updated_at"`
}
