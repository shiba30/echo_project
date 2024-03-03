package model

import (
	"time"
)

type User struct {
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Email            string    `gorm:"size:255;not null;unique" json:"email"`
	Password         string    `gorm:"size:255;not null" json:"-"`
	PhoneNumber      string    `gorm:"size:255;unique" json:"phone_number"`
	DisplayName      string    `gorm:"size:50;not null" json:"display_name"`
	SelfIntroduction string    `gorm:"type:text" json:"self_introduction"`
	Location         string    `gorm:"size:255" json:"location"`
	Website          string    `gorm:"size:255" json:"website"`
	BirthDate        time.Time `json:"birth_date"`
	ProfileImage     string    `gorm:"size:255" json:"profile_image"`
	AvatarImage      string    `gorm:"size:255" json:"avatar_image"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ActivationToken  string    `gorm:"size:255" json:"activation_token"`
	IsActive         bool      `gorm:"default:false" json:"is_active"`
}

// TableName GORMにテーブル名を指定するためのメソッド
func (User) TableName() string {
	return "users"
}
