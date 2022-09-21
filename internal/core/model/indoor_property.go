package model

import "time"

type IndoorProperty struct {
	Time    time.Time `gorm:"column:created_at" json:"created_at"`
	Type    string    `gorm:"column:type" json:"type"`
	Value   float32   `gorm:"column:value" json:"value"`
	Machine string    `gorm:"column:machine" json:"machine_id"`
}
