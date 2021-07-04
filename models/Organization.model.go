package models

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Location string    `json:"location"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Timezone time.Time `json:"timezone"`
	Status   string    `json:"status"`
}
