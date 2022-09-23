package service

import (
	"cpfd-back/internal/api/repo"
	g "cpfd-back/internal/conf/grpc"
	"cpfd-back/internal/core/model"
)

type CreateParticleLogParams struct {
	Pm1        float32 `json:"pm1"`
	Pm25       float32 `json:"pm2_5"`
	Pm10       float32 `json:"pm10"`
	MachineNum int     `json:"machine"`
}

type CreateIndoorPropertyLogParams struct {
	Type       string  `json:"type"`
	Value      float32 `json:"value"`
	MachineNum int     `json:"machine_num"`
}

type ServiceManager interface {
	CreateMachine(p repo.CreateMachineParams) error
	GetMachine(machineUid string, machineNum int) (*model.Machine, error)
	GetMachineNum(machineUid string) (int, error)

	GetAllLogsToCSV() (string, error)

	CreateActivityLog(name string, action bool) error
	GetActivityLogs() ([]model.Activity, error)
	GetActivityLogsToCSV() (string, error)
	GetCurrentActivity() (map[string]bool, error)

	CreateParticleLog(p CreateParticleLogParams) error
	GetAllParticleLogs() ([]model.Particle, error)
	GetAllParticleLogsToCSV() (string, error)
	GetParticleLogsWithDates(p repo.GetParticleLogsWithDatesParams) ([]model.Particle, error)
	GetParticleLogsWithDatesToCSV(p repo.GetParticleLogsWithDatesParams) (string, error)

	CreateIndoorPropertyLog(p CreateIndoorPropertyLogParams) error
	GetIndoorPropertyLogs(p repo.GetIndoorPropertyLogsParams) ([]model.IndoorProperty, error)
	GetIndoorPropertyLogsToCSV(p repo.GetIndoorPropertyLogsParams) (string, error)
}

func New(r repo.RepoManager, dg *g.DataGen) ServiceManager {
	return &Service{
		repo:          r,
		dataGenerator: dg,
	}
}

type Service struct {
	repo          repo.RepoManager
	dataGenerator *g.DataGen
}
