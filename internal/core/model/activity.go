package model

import "time"

const ActivityTable = "activities"

type Activity struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Name      string    `gorm:"column:name" json:"name"`
	Action    bool      `gorm:"column:action" json:"action"`
	Type      string    `gorm:"column:type" json:"type"`
}
