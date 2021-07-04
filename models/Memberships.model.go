package models

import (
	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	UserID         int          `json:"userid"`
	User           User         `json:"user"`
	OrganizationID int          `json:"orgid"`
	Organization   Organization `json:"organization"`
	Name           string       `json:"name"`
	Role           string       `json:"role"`
}
