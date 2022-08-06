package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
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

func (s *IndoorPropertyService) GetLogToCSV(startTime, endTime time.Time) (string, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	if start == "0001-01-01 00:00:00" {
		start = core.StartDate
	}
	if end == "0001-01-01 00:00:00" {
		end = time.Now().Format("2006-01-02 15:04:05")
	}

	return s.Repo.GetLogToCSV(start, end)
}
