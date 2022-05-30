package repo

import (
	"cpfd-back/internal/core/model"
	"gorm.io/gorm"
	"log"
)

type IndoorPropertyRepo struct {
	Mysql *gorm.DB
}

func NewIndoorPropertyRepo(mysqlDb *gorm.DB) *IndoorPropertyRepo {
	return &IndoorPropertyRepo{
		Mysql: mysqlDb,
	}
}

func (r *IndoorPropertyRepo) CreateLog(ip model.IndoorProperty) error {
	if err := r.Mysql.Create(&ip).Error; err != nil {
		log.Printf("[ERROR] Failed to create indoor property log: %v\n", err)
		return err
	}
	return nil
}
