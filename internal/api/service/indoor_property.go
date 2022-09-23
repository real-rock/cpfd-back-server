package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"cpfd-back/internal/util"
	"time"
)

func (s *Service) CreateIndoorPropertyLog(p CreateIndoorPropertyLogParams) error {
	machine, err := s.repo.GetMachineWithNum(p.MachineNum)
	if err != nil {
		return err
	}
	param := repo.CreateIndoorPropertyLogParams{
		Type:      p.Type,
		Value:     p.Value,
		MachineID: machine.Id.String(),
	}
	return s.repo.CreateIndoorPropertyLog(param)
}

func (s *Service) GetIndoorPropertyLogs(p repo.GetIndoorPropertyLogsParams) ([]model.IndoorProperty, error) {
	if p.Start.IsZero() && p.End.IsZero() {
		return s.repo.GetAllIndoorPropertyLogs()
	} else if p.Start.IsZero() {
		st, _ := time.Parse(core.TimeFormat, core.StartDate)
		p.Start = st
	}
	return s.repo.GetIndoorPropertyLogsWithDates(p)
}

func (s *Service) GetIndoorPropertyLogsToCSV(p repo.GetIndoorPropertyLogsParams) (string, error) {
	var ips interface{}

	if p.Start.IsZero() && p.End.IsZero() {
		ips, err := s.repo.GetAllIndoorPropertyLogs()
		if err != nil {
			return "", err
		}
		return util.WriteTempCSV(ips, "ip_")
	} else if p.Start.IsZero() {
		st, _ := time.Parse(core.TimeFormat, core.StartDate)
		p.Start = st
	}
	ips, err := s.repo.GetIndoorPropertyLogsWithDates(p)
	if err != nil {
		return "", err
	}
	return util.WriteTempCSV(ips, "ip_")
}
