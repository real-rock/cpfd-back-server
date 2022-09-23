package service

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"

	"github.com/google/uuid"
)

func (s *Service) CreateMachine(p repo.CreateMachineParams) error {
	return s.repo.CreateMachine(p)
}

func (s *Service) GetMachine(machineUid string, machineNum int) (*model.Machine, error) {
	if machineUid == "" {
		return s.repo.GetMachineWithNum(machineNum)
	} else {
		uid, err := uuid.FromBytes([]byte(machineUid))
		if err != nil {
			log.Logger.Errorln("failed to get uuid from request: ", err.Error())
			return nil, err
		}
		return s.repo.GetMachine(uid)
	}
}

func (s *Service) GetMachineNum(machineUid string) (int, error) {
	uid, err := uuid.FromBytes([]byte(machineUid))
	if err != nil {
		log.Logger.Errorln("failed to get uuid from request: ", err.Error())
		return 0, err
	}
	return s.repo.GetMachineNum(uid)
}
