package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	Mfa      bool   `json:"mfa"`
	Mfa_type string `json:"mfa_type"`
}
