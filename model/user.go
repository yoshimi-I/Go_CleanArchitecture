package model

import (
	"time"
)

// user型を指定
type User struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Email		 string   `json:"email" gorm:"unique"`
	Password string   `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 返却の型を指定
type UserResponse struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Email		 string   `json:"email" gorm:"unique"`
}
