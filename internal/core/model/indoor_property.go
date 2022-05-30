package model

import "time"

type IndoorProperty struct {
	Time    time.Time `gorm:"column:time" json:"time"`
	Type    string    `gorm:"column:type" json:"type"`
	Value   float32   `gorm:"column:value" json:"value"`
	Machine int       `gorm:"column:machine" json:"machine"`
}
