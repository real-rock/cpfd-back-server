package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"time"
)

type ParticleService struct {
	repo *repo.ParticleRepo
}

func NewParticleService(repo *repo.ParticleRepo) *ParticleService {
	return &ParticleService{
		repo: repo,
	}
}

func (s *ParticleService) GetLogs() ([]model.Particle, error) {
	return s.repo.GetAllLogs()
}

func (s *ParticleService) GetLogToFile(startTime, endTime time.Time) (string, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	return s.repo.GetLogsToFile(start, end)
}

func (s *ParticleService) GetChartData(startTime, endTime time.Time) (map[string][]map[string]interface{}, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	return s.repo.GetLogsWithDates(start, end)
}

func (s *ParticleService) CreateLog(p model.Particle) error {
	p.Time = time.Now().In(core.Location)
	return s.repo.CreateLog(p)
}
