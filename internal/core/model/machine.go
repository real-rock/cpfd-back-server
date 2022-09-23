package model

import "github.com/google/uuid"

type Machine struct {
	Id          uuid.UUID `gorm:"column:id"`
	Description string    `gorm:"column:description"`
	MachineNum  int       `gorm:"column:machine_num"`
	Location    string    `gorm:"column:location"`
	PostalCode  string    `gorm:"column:postal_code"`
}
