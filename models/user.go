package models

type Role string

const (
	CustomerRole Role = "customer"
	MerchantRole Role = "merchant"
)

type User struct {
	Id       uint      `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password []byte    `json:"-"`
	Role     Role      `json:"role" gorm:"default:customer"`
	Payments []Payment `gorm:"foreignKey:UserID" json:"-"`
}
