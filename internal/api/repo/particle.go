package repo

import (
	"cpfd-back/internal/core/model"
	"fmt"
	"gorm.io/gorm"
	"log"
)

const dateFormat = "%Y-%m-%d %H:%i"

type ParticleRepo struct {
	Mysql *gorm.DB
}

func NewParticleRepo(sqlDb *gorm.DB) *ParticleRepo {
	return &ParticleRepo{
		Mysql: sqlDb,
	}
}

func (r *ParticleRepo) GetAllLogs() ([]model.Particle, error) {
	var particles []model.Particle

	if err := r.Mysql.Find(&particles).Error; err != nil {
		log.Printf("[ERROR] Failed to get particles from sql db: %v", err)
		return nil, err
	}
	return particles, nil
}

func (r *ParticleRepo) GetLogsWithDates(start, end string) (map[string][]map[string]interface{}, error) {
	particles := make(map[string][]map[string]interface{})
	machines := []string{"107", "120", "121", "124", "134", "181", "196", "199"}

	for _, machine := range machines {
		var particle []map[string]interface{}

		sql := fmt.Sprintf("select UNIX_TIMESTAMP(time) as time, pm1, pm2_5, pm10 from particles where machine='%s' and time between '%s' and '%s' order by time",
			machine, start, end)

		if err := r.Mysql.Raw(sql).Scan(&particle).Error; err != nil {
			log.Printf("[ERROR] Failed to get particles: %v", err)
			return nil, err
		}
		particles[machine] = particle
	}
	return particles, nil
}

func (r *ParticleRepo) CreateLog(p model.Particle) error {
	return r.Mysql.Create(&p).Error
}
