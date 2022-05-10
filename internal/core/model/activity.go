package model

import "time"

const ActivityTable = "activities"

type Activity struct {
	Name      string    `gorm:"column:name" json:"name"`
	TimeStamp time.Time `gorm:"column:time" json:"-"`
	Action    bool      `gorm:"column:action" json:"action"`
	Type      string    `gorm:"column:type" json:"type"`
}
