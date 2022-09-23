package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"cpfd-back/internal/util"

	"golang.org/x/exp/slices"
)

func (s *Service) CreateActivityLog(name string, action bool) error {
	var p repo.CreateActivityLogParams

	if slices.Contains(core.Members, name) {
		p.Type = "PERSON"
	} else {
		p.Type = name
	}
	p.Name = name
	p.Action = action
	return s.repo.CreateActivityLog(p)
}

func (s *Service) GetActivityLogs() ([]model.Activity, error) {
	return s.repo.GetActivityLogs()
}

func (s *Service) GetActivityLogsToCSV() (string, error) {
	activities, err := s.GetActivityLogs()
	if err != nil {
		return "", err
	}
	name, err := util.WriteTempCSV(activities, "activities_")
	if err != nil {
		return "", err
	}
	return name, nil
}

func (s *Service) GetCurrentActivity() (map[string]bool, error) {
	return s.repo.GetCurrentActivity()
}
