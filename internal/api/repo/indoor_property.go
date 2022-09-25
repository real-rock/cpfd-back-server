package repo

import (
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"
)

func (r *Repo) CreateIndoorPropertyLog(p CreateIndoorPropertyLogParams) error {
	if err := r.Mysql.Table("indoor_properties").Create(&p).Error; err != nil {
		log.Logger.Errorln("failed to create indoor property log: ", err)
		return err
	}
	return nil
}

func (r *Repo) GetAllIndoorPropertyLogs() ([]model.IndoorProperty, error) {
	var res []model.IndoorProperty

	sql := `SELECT created_at, type, value, machine, machine_num
	FROM indoor_properties as ip LEFT JOIN machines m ON m.id = ip.machine;`

	if err := r.Mysql.Raw(sql).Scan(&res).Error; err != nil {
		log.Logger.Errorln("failed to fetch indoor-properties logs: ", err)
		return nil, err
	}
	return res, nil
}

func (r *Repo) GetIndoorPropertyLogsWithDates(p GetIndoorPropertyLogsParams) ([]model.IndoorProperty, error) {
	var res []model.IndoorProperty

	sql := `SELECT created_at, type, value, machine, machine_num
	FROM indoor_properties as ip LEFT JOIN machines m ON m.id = ip.machine
	WHERE created_at BETWEEN ? AND ?;`

	if err := r.Mysql.Raw(sql, p.Start, p.End).Scan(&res).Error; err != nil {
		log.Logger.Errorln("failed to fetch indoor-properties logs: ", err)
		return nil, err
	}
	return res, nil
}
