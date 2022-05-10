package repo

import (
	"context"
	"cpfd-back/internal/core"
	"cpfd-back/internal/core/model"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"time"
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
		sqlCmd := fmt.Sprintf("select action from %s where `name` = '%s' order by time DESC limit 1;",
			model.ActivityTable, member)

		if err := r.Mysql.Raw(sqlCmd).Scan(&res).Error; err != nil {
			panic(err)
		}
		log.Printf("[INFO] action successfully fetched")
		if err := r.Redis.Set(c, member, res.Action, 0).Err(); err != nil {
			panic(err)
		}
		log.Printf("[INFO] {key: %s, val: %v} successfully saved to redis db", member, res.Action)
	}
}

func (r *InoutRepo) GetLogs() ([][]string, error) {
	var activities []map[string]interface{}

	if err := r.Mysql.Table(model.ActivityTable).Find(&activities).Error; err != nil {
		log.Printf("[ERROR] Failed to get logs from mysql: %v", err)
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
			log.Printf("[WARNING] Can't find key %s, return false default", member)
			action = "false"
		} else if err != nil {
			log.Printf("[ERROR] Failed fetching value from key %s: %v", member, err)
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
		Name:      name,
		TimeStamp: time.Now().In(core.Location),
		Action:    action,
		Type:      objType,
	}
	if err := r.Mysql.Create(&activity).Error; err != nil {
		log.Printf("[ERROR] error while saving data: %v", err)
		return err
	}
	if err := r.Redis.Set(context.Background(), name, action, 0).Err(); err != nil {
		log.Printf("[ERROR] failed to save information to redis db: %v", err)
		return err
	}
	return nil
}
