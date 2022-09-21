package repo

import (
	"context"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	log "cpfd-back/internal/log"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type InoutRepo struct {
	Mysql *gorm.DB
	Redis *redis.Client
}

func NewInoutRepo(mysqlDb *gorm.DB, redisDb *redis.Client) *InoutRepo {
	return &InoutRepo{
		Mysql: mysqlDb,
		Redis: redisDb,
	}
}

func (r *InoutRepo) Init() {
	c := context.Background()
	var res struct {
		Action bool `gorm:"column:action"`
	}

	for _, member := range core.Objects {
		sqlCmd := fmt.Sprintf("SELECT action FROM %s WHERE `name` = '%s' ORDER BY created_at DESC LIMIT 1;",
			model.ActivityTable, member)

		if err := r.Mysql.Raw(sqlCmd).Scan(&res).Error; err != nil {
			panic(err)
		}
		log.Logger.Infoln("action successfully fetched")
		if err := r.Redis.Set(c, member, res.Action, 0).Err(); err != nil {
			panic(err)
		}
		log.Logger.Infof("{key: %s, val: %v} successfully saved to redis db\n", member, res.Action)
	}
}

func (r *InoutRepo) GetLogs() ([][]string, error) {
	var activities []map[string]interface{}

	if err := r.Mysql.Table(model.ActivityTable).Limit(100).Order("time desc").Find(&activities).Error; err != nil {
		log.Logger.Errorln("failed to get logs from mysql: ", err.Error())
		return nil, err
	}
	values := make([][]string, len(activities))
	i := 0
	for _, activity := range activities {
		vals := make([]string, 4)

		vals[0] = fmt.Sprintf("%v", activity["name"])
		vals[1] = activity["time"].(time.Time).Format("2006-01-02 15:04:05")
		vals[2] = fmt.Sprintf("%v", activity["action"])
		vals[3] = fmt.Sprintf("%v", activity["type"])

		values[i] = vals
		i += 1
	}
	return values, nil
}

func (r *InoutRepo) GetCurrentInfo() (map[string]bool, error) {
	m := make(map[string]bool)
	c := context.Background()

	for _, member := range core.Objects {
		action, err := r.Redis.Get(c, member).Result()
		if err == redis.Nil {
			log.Logger.Warnf("can't find key %s, return false default\n", member)
			action = "false"
		} else if err != nil {
			log.Logger.Errorf("failed fetching value from key %s: %v\n", member, err)
			return nil, err
		}
		if action == "1" {
			m[member] = true
		} else {
			m[member] = false
		}
	}
	return m, nil
}

func (r *InoutRepo) CreateLog(name, objType string, action bool) error {
	activity := model.Activity{
		Name: name,
		//TimeStamp: time.Now().In(core.Location),
		TimeStamp: time.Now(),
		Action:    action,
		Type:      objType,
	}
	if err := r.Mysql.Create(&activity).Error; err != nil {
		log.Logger.Errorln("error while saving data: ", err.Error())
		return err
	}
	if err := r.Redis.Set(context.Background(), name, action, 0).Err(); err != nil {
		log.Logger.Errorln("failed to save information to redis db: ", err.Error())
		return err
	}
	return nil
}
