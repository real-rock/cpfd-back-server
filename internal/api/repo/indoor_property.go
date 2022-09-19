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

type IndoorPropertyRepo struct {
	Mysql *gorm.DB
}

func NewIndoorPropertyRepo(mysqlDb *gorm.DB) *IndoorPropertyRepo {
	return &IndoorPropertyRepo{
		Mysql: mysqlDb,
	}
}

func (r *IndoorPropertyRepo) CreateLog(ip model.IndoorProperty) error {
	if err := r.Mysql.Create(&ip).Error; err != nil {
		log.Logger.Errorln("failed to create indoor property log: ", err)
		return err
	}
	return nil
}

func (r *IndoorPropertyRepo) GetLogToCSV(start, end string) (string, error) {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	name := "ip_" + strconv.Itoa(num)

	path := core.MysqlFilePath + "/" + name + ".csv"

	sql := fmt.Sprintf("select 'DATE', 'TYPE', 'VALUE', 'MACHINE' union all "+
		"select time, type, value, machine from indoor_properties where time between '%s' and '%s' "+
		"into outfile '%s' FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'", start, end, path)

	if err := r.Mysql.Exec(sql).Error; err != nil {
		log.Logger.Errorln("failed to create indoor-properties file: ", err)
		return "", err
	}
	return core.FileDir + "/" + name + ".csv", nil
}
