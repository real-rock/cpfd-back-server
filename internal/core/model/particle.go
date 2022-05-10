package model

import "time"

type Particle struct {
	Time    time.Time `gorm:"column:time" json:"time"`
	Pm1     float32   `gorm:"column:pm1" json:"pm1"`
	Pm25    float32   `gorm:"column:pm2_5" json:"pm2_5"`
	Pm10    float32   `gorm:"column:pm10" json:"pm10"`
	Machine string    `gorm:"column:machine" json:"machine"`
}
