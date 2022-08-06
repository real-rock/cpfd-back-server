package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
	"encoding/csv"
	"golang.org/x/exp/slices"
	"log"
	"os"
)

type InoutService struct {
	Repo *repo.InoutRepo
}

func NewInoutService(repo *repo.InoutRepo) *InoutService {
	s := InoutService{
		Repo: repo,
	}
	s.Init()
	return &s
}

func (s *InoutService) Init() {
	s.Repo.Init()
}

func (s *InoutService) GetAllLogsToFiles() ([][]string, error) {
	rows, err := s.Repo.GetLogs()
	if err != nil {
		return nil, err
	}
	f, err := os.Create("test.csv")
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("[ERROR] Failed to close csv file: %v", err)
			return
		}
	}()

	if err != nil {
		log.Printf("[ERROR] Failed to open csv file: %v", err)
		return nil, err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(core.CsvHeaders); err != nil {
		log.Printf("[ERROR] Failed to write header to file: %v", err)
		return nil, err
	}
	if err := w.WriteAll(rows); err != nil {
		log.Printf("[ERROR] Failed to write file: %v", err)
		return nil, err
	}
	return rows, nil
}

func (s *InoutService) GetLogs() ([][]string, error) {
	return s.Repo.GetLogs()
}

func (s *InoutService) GetCurrentInfo() (map[string]bool, error) {
	return s.Repo.GetCurrentInfo()
}

func (s *InoutService) CreateLog(name string, action bool) error {
	var objType string

	if slices.Contains(core.Members, name) == true {
		objType = "PERSON"
	} else {
		objType = name
	}
	return s.Repo.CreateLog(name, objType, action)
}
