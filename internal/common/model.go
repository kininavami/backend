package common

import "time"

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time		`json:"update_at"`
}
