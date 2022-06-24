package entities

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID         int       `gorm:"primary_key;auto_increment" json:"id"`
	Body       string    `gorm:"type:text;not_null" json:"body"`
	IsJoinable bool      `gorm:"type:boolean;not_null" json:"is_joinable"`
	Owner      User      `gorm:"foreignkey:OwnerID" json:"owner"`
	OwnerID    uuid.UUID `gorm:"type:uuid;not_null" json:"owner_id"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp;not_null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp;not_null" json:"updated_at"`
}
