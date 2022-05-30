package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core/model"
	"time"
)

type IndoorPropertyService struct {
	Repo *repo.IndoorPropertyRepo
}

func NewIndoorPropertyService(repo *repo.IndoorPropertyRepo) *IndoorPropertyService {
	return &IndoorPropertyService{
		Repo: repo,
	}
}

func (s *IndoorPropertyService) CreateLog(ip model.IndoorProperty) error {
	ip.Time = time.Now()
	return s.Repo.CreateLog(ip)
}
