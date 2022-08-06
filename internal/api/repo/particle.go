package repo

import (
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"time"
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

func (r *ParticleRepo) GetLogsToFile(start, end string) ([]string, error) {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	name := "particle_" + strconv.Itoa(num)
	activityName := "activity_" + strconv.Itoa(num)

	particlePath := core.MysqlFilePath + "/" + name + ".csv"
	activityPath := core.MysqlFilePath + "/" + activityName + ".csv"

	sql := fmt.Sprintf("select DATE_FORMAT(time, '%s'), pm1, pm2_5, pm10, machine "+
		"from particles where time between '%s' and '%s' "+
		"into outfile '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", dateFormat, start, end, particlePath)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Printf("[ERROR] Failed to create particle file: %v", err)
		return nil, err
	}
	sql = fmt.Sprintf("select name, DATE_FORMAT(time, '%s'), action, type "+
		"from activities where time between '%s' and '%s' "+
		"into outfile '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", dateFormat, start, end, activityPath)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Printf("[ERROR] Failed to create activity file: %v", err)
		return nil, err
	}
	return []string{core.FileDir + "/" + name + ".csv", core.FileDir + "/" + activityName + ".csv"}, nil
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
