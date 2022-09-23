package repo

import (
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"

	"github.com/google/uuid"
)

func (r *Repo) CreateMachine(p CreateMachineParams) error {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Logger.Errorln("failed to create uuid: ", err.Error())
		return err
	}
	machine := model.Machine{
		Id:          newUUID,
		Description: p.Description,
		Location:    p.Location,
		MachineNum:  p.MachineNum,
		PostalCode:  p.PostalCode,
	}
	if err := r.Mysql.Create(&machine).Error; err != nil {
		log.Logger.Errorln("failed to create machine: ", err.Error())
		return err
	}
	return nil
}

func (r *Repo) GetMachine(machineUid uuid.UUID) (*model.Machine, error) {
	machine := model.Machine{
		Id: machineUid,
	}
	if err := r.Mysql.First(&machine).Error; err != nil {
		log.Logger.Errorln("failed to get machine: ", err.Error())
		return nil, err
	}
	return &machine, nil
}

func (r *Repo) GetMachineWithNum(num int) (*model.Machine, error) {
	var machine model.Machine
	if err := r.Mysql.First(&machine, "machine_num = ?", num).Error; err != nil {
		log.Logger.Errorln("failed to get machine: ", err.Error())
		return nil, err
	}
	return &machine, nil
}

func (r *Repo) GetMachineNum(machineUid uuid.UUID) (int, error) {
	machine, err := r.GetMachine(machineUid)
	if err != nil {
		return -1, err
	}
	return machine.MachineNum, nil
}
