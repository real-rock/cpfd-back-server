package model

import "time"

type Particle struct {
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	Pm1        float32   `gorm:"column:pm1" json:"pm1"`
	Pm25       float32   `gorm:"column:pm2_5" json:"pm2_5"`
	Pm10       float32   `gorm:"column:pm10" json:"pm10"`
	MachineID  string    `gorm:"column:machine" json:"machine_id"`
	MachineNum int       `gorm:"column:machine_num" json:"machine_num"`
}
