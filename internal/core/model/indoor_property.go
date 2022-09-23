package model

import (
	"time"
)

type IndoorProperty struct {
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	Type       string    `gorm:"column:type" json:"type"`
	Value      float32   `gorm:"column:value" json:"value"`
	MachineID  string    `gorm:"column:machine" json:"machine_id"`
	MachineNum int       `gorm:"column:machine_num" json:"machine_num"`
}
