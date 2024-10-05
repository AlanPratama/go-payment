package models

import "time"

type Status string

const (
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

type Type string

const (
	LoginType  Type = "login"
	LogoutType Type = "logout"
)

type History struct {
	Id        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Type      Type      `json:"type" gorm:"type:enum('login','logout')"`
	Status    Status    `json:"status" gorm:"type:enum('success','failed')"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
