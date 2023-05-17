package model

import "time"

type Task struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Title     string   `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User 		User     `json:"user" gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // User削除に伴いそれに付随するTaskも削除
	UserId    uint     `json:"user_id" gorm:"not null"`
}

// Taskの返却型
type TaskResponse struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Title     string   `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

