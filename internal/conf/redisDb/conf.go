package redisDb

import (
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

type Redis struct {
	conf
	DB *redis.Client
}

func New() *Redis {
	r := Redis{}
	r.conf = newConf()
	r.Conn()
	r.Info()
	return &r
}

func (r *Redis) Conn() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: r.pwd,
		DB:       r.db,
	})
	r.DB = rdb
}
