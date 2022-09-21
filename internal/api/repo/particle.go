package repo

import (
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

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
		log.Logger.Errorf("failed to get particles from sql db: %v", err)
		return nil, err
	}
	return particles, nil
}

func (r *ParticleRepo) GetAllLogsToFile(start, end string) ([]string, error) {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	name := "particle_" + strconv.Itoa(num)
	activityName := "activity_" + strconv.Itoa(num)

	particlePath := core.MysqlFilePath + "/" + name + ".csv"
	activityPath := core.MysqlFilePath + "/" + activityName + ".csv"

	sql := fmt.Sprintf("SELECT created_at, pm1, pm2_5, pm10, machine_num "+
		"FROM particles left join machines m on m.id = particles.machine"+
		"WHERE created_at BETWEEN '%s' AND '%s' "+
		"INTO OUTFILE '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", start, end, particlePath)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Logger.Errorf("failed to create particle file: %v", err)
		return nil, err
	}
	sql = fmt.Sprintf("SELECT name, created_at, action, type "+
		"FROM activities WHERE created_at BETWEEN '%s' AND '%s' "+
		"INTO OUTFILE '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", start, end, activityPath)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Logger.Errorf("failed to create activity file: %v", err)
		return nil, err
	}
	return []string{core.FileDir + "/" + name + ".csv", core.FileDir + "/" + activityName + ".csv"}, nil
}

func (r *ParticleRepo) GetLogsWithDates(start, end string) (map[string][]map[string]interface{}, error) {
	particles := make(map[string][]map[string]interface{})
	machines := []string{"107", "120", "121", "124", "134", "181", "196", "199"}

	for _, machine := range machines {
		var particle []map[string]interface{}

		sql := fmt.Sprintf("SELECT UNIX_TIMESTAMP(created_at) as time, pm1, pm2_5, pm10 "+
			"FROM particles WHERE machine='%s' AND time BETWEEN '%s' AND '%s' ORDER BY time",
			machine, start, end)

		if err := r.Mysql.Raw(sql).Scan(&particle).Error; err != nil {
			log.Logger.Errorf("failed to get particles: %v", err)
			return nil, err
		}
		particles[machine] = particle
	}
	return particles, nil
}

func (r *ParticleRepo) GetLogsToFile(machine []string, start, end string) (string, error) {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	name := "particle_" + strconv.Itoa(num)

	particlePath := core.MysqlFilePath + "/" + name + ".csv"
	machineStr := ""
	for _, m := range machine {
		machineStr = machineStr + m
	}

	sql := fmt.Sprintf("SELECT 'DATE', 'PM1', 'PM2.5', 'PM10', 'MACHINE' UNION ALL "+
		"SELECT created_at, pm1, pm2_5, pm10, machine_num "+
		"FROM particles left join machines m on m.id = particles.machine "+
		"WHERE machine in (%s) and TIME BETWEEN '%s' AND '%s' "+
		"INTO OUTFILE '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", machineStr, start, end, particlePath)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Logger.Errorf("failed to create particle file: %v", err)
		return "", err
	}
	return core.FileDir + "/" + name + ".csv", nil
}

func (r *ParticleRepo) CreateLog(p model.Particle) error {
	return r.Mysql.Create(&p).Error
}
