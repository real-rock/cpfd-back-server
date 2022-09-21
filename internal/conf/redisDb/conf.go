package redisDb

import (
	"fmt"
	"os"
	"strconv"
)

type conf struct {
	addr string
	pwd  string
	db   int
	port string
}

func newConf() conf {
	addr := os.Getenv("REDIS_ADDR")
	pwd := os.Getenv("REDIS_PWD")
	db := os.Getenv("REDIS_DB")
	port := os.Getenv("REDIS_PORT")

	intDb, err := strconv.Atoi(db)
	if err != nil {
		panic(err)
	}
	return conf{
		addr: addr,
		pwd:  pwd,
		db:   intDb,
		port: port,
	}
}

func (c conf) Info() {
	fmt.Println("========== REDIS ==========")
	fmt.Println("Address: ", c.addr)
	fmt.Println("Password: ", c.pwd)
	fmt.Println("DB: ", c.db)
	fmt.Println("Port: ", c.port)
}
