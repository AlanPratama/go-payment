package models

import "time"

type PaymentStatus string

const (
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)

type Payment struct {
	Id            uint          `json:"id"`
	Amount        uint          `json:"amount"`
	UserID        uint          `json:"user_id"`
	TargetUserID  uint          `json:"target_user_id"`
	PaymentStatus PaymentStatus `json:"payment_status" gorm:"type:enum('success','failed', 'pending')"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
