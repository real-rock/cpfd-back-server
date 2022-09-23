package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"cpfd-back/internal/util"
	"time"
)

func (s *Service) CreateParticleLog(p CreateParticleLogParams) error {
	machine, err := s.repo.GetMachineWithNum(p.MachineNum)
	if err != nil {
		return err
	}
	param := repo.CreateParticleLogParams{
		Pm1:       p.Pm1,
		Pm25:      p.Pm25,
		Pm10:      p.Pm10,
		MachineID: machine.Id.String(),
	}
	return s.repo.CreateParticleLog(param)
}

func (s *Service) GetAllParticleLogs() ([]model.Particle, error) {
	return s.repo.GetAllParticleLogs()
}

func (s *Service) GetAllParticleLogsToCSV() (string, error) {
	particles, err := s.repo.GetAllParticleLogs()
	if err != nil {
		return "", nil
	}
	return util.WriteTempCSV(particles, "particle_")
}

func (s *Service) GetParticleLogsWithDates(p repo.GetParticleLogsWithDatesParams) ([]model.Particle, error) {
	if p.Start.IsZero() && p.End.IsZero() {
		return s.repo.GetAllParticleLogs()
	} else if p.Start.IsZero() {
		st, _ := time.Parse(core.TimeFormat, core.StartDate)
		p.Start = st
	}
	return s.repo.GetParticleLogsWithDates(p)
}

func (s *Service) GetParticleLogsWithDatesToCSV(p repo.GetParticleLogsWithDatesParams) (string, error) {
	particles, err := s.repo.GetParticleLogsWithDates(p)
	if err != nil {
		return "", nil
	}
	return util.WriteTempCSV(particles, "particle_d_")
}
