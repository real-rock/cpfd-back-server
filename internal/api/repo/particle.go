package repo

import (
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"
)

func (r *Repo) CreateParticleLog(p CreateParticleLogParams) error {
	if err := r.Mysql.Table("particles").Create(&p).Error; err != nil {
		log.Logger.Errorln("failed to create particle: ", err.Error())
		return err
	}
	return nil
}

func (r *Repo) GetAllParticleLogs() ([]model.Particle, error) {
	var particles []model.Particle

	cmd := `SELECT created_at, pm1, pm2_5, pm10, machine, machine_num
	FROM particles left join machines m on m.id = particles.machine
	ORDER BY created_at;`

	if err := r.Mysql.Raw(cmd).Scan(&particles).Error; err != nil {
		log.Logger.Errorf("failed to get particles from sql db: %v", err)
		return nil, err
	}
	return particles, nil
}

func (r *Repo) GetParticleLogsWithDates(p GetParticleLogsWithDatesParams) ([]model.Particle, error) {
	var particles []model.Particle

	cmd := `SELECT created_at, pm1, pm2_5, pm10, machine, machine_num 
	FROM particles left join machines m on m.id = particles.machine
	WHERE created_at BETWEEN ? and ?
	ORDER BY created_at;`

	if err := r.Mysql.Raw(cmd, p.Start, p.End).Scan(&particles).Error; err != nil {
		log.Logger.Errorf("failed to get particles from sql db: %v", err)
		return nil, err
	}
	return particles, nil
}
