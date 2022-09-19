package service

import (
	"cpfd-back/internal/api/repo"
	g "cpfd-back/internal/conf/grpc"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"cpfd-back/internal/core/pb"
	"log"
	"os"
	"time"
)

type ParticleService struct {
	repo *repo.ParticleRepo
	grpc *g.DataGen
}

func NewParticleService(repo *repo.ParticleRepo, grpc *g.DataGen) *ParticleService {
	return &ParticleService{
		repo: repo,
		grpc: grpc,
	}
}

func (s *ParticleService) GetLogs() ([]model.Particle, error) {
	return s.repo.GetAllLogs()
}

func (s *ParticleService) GetAllLogToFile(startTime, endTime time.Time, method string) (string, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	if start == "0001-01-01 00:00:00" {
		start = core.StartDate
	}
	if end == "0001-01-01 00:00:00" {
		end = time.Now().Format("2006-01-02 15:04:05")
	}
	if method == "" {
		method = "mean"
	}

	paths, err := s.repo.GetAllLogsToFile(start, end)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := os.Remove(paths[0]); err != nil {
			log.Println(err)
		}
		if err := os.Remove(paths[1]); err != nil {
			log.Println(err)
		}
	}()

	req := pb.ParticleReq{
		ParticlePath: paths[0],
		ActivityPath: paths[1],
		Start:        start,
		End:          end,
		Method:       method,
	}
	res, err := s.grpc.Request(&req)
	if err != nil {
		log.Printf("[ERROR] Failed to get response from grpc server: %v", err)
		return "", err
	}
	return res.FilePath, nil
}

func (s *ParticleService) GetLogToFile(machines []string, startTime, endTime time.Time) (string, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	if start == "0001-01-01 00:00:00" {
		start = core.StartDate
	}
	if end == "0001-01-01 00:00:00" {
		end = time.Now().Format("2006-01-02 15:04:05")
	}

	return s.repo.GetLogsToFile(machines, start, end)
}

func (s *ParticleService) GetChartData(startTime, endTime time.Time) (map[string][]map[string]interface{}, error) {
	start := startTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")

	return s.repo.GetLogsWithDates(start, end)
}

func (s *ParticleService) CreateLog(p model.Particle) error {
	//p.Time = time.Now().In(core.Location)
	p.Time = time.Now()
	return s.repo.CreateLog(p)
}
