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

	viewSql := fmt.Sprintf("create view %s as(select DATE_FORMAT(p.time, '%s') as f_time, AVG(p.pm1) as avg_pm1, AVG(p.pm2_5) as avg_pm2_5, AVG(p.pm10) as avg_pm10, m.location "+
		"from particles as p join machines m on m.id = p.machine "+
		"where time between '%s' and '%s' "+
		"group by f_time, location);", name, dateFormat, start, end)

	if err := r.Mysql.Exec(viewSql).Error; err != nil {
		log.Printf("[ERROR] Failed to create view '%s': %v", start+end, err)
		return nil, err
	}
	defer func() {
		dropViewSql := fmt.Sprintf("drop view %s;", name)
		if err := r.Mysql.Exec(dropViewSql).Error; err != nil {
			log.Printf("[ERROR] Failed to drop view '%v': %v", name, err)
			return
		}
	}()
	sql := fmt.Sprintf("select 'DATE', 'PM1', 'PM2.5', 'PM10', 'PM1_OUT', 'PM2.5_OUT', 'PM10_OUT', 'PM1_H_OUT', 'PM2.5_H_OUT', 'PM10_H_OUT' union all"+
		" select tt.time, tt.pm1, tt.pm2_5, tt.pm10, gg.pm1_out, "+
		"gg.pm2_5_out, gg.pm10_out, vv.pm1, vv.pm2_5, vv.pm10 into outfile '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n' from ("+
		"select f_time as time, avg_pm1 as pm1, avg_pm2_5 as pm2_5, avg_pm10 as pm10 from %s where location = 'IN') as tt "+
		"left outer join (select f_time as time, avg_pm1 as pm1_out, avg_pm2_5 as pm2_5_out, avg_pm10 as pm10_out from %s where location = 'OUT') as gg on tt.time = gg.time left outer join ("+
		"select f_time as time, avg_pm1 as pm1, avg_pm2_5 as pm2_5, avg_pm10 as pm10 from %s where location = 'HALL_OUT') as vv on tt.time = vv.time;", particlePath, name, name, name)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Printf("[ERROR] Failed to generate data: %v", err)
		return nil, err
	}

	activitySql := fmt.Sprintf("select 'NAME', 'TIME', 'ACTION', 'TYPE' union all "+
		"select * into outfile '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n' from activities where time < '%s'", activityPath, end)
	if err := r.Mysql.Exec(activitySql).Error; err != nil {
		log.Printf("[ERROR] Failed to create activity file")
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
