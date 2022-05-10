package redisDb

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type conf struct {
	addr string
	pwd  string
	db   int
}

func newConf() conf {
	log.SetFlags(log.Ltime | log.Lshortfile)

	addr := os.Getenv("REDIS_ADDR")
	pwd := os.Getenv("REDIS_PWD")
	db := os.Getenv("REDIS_DB")

	if addr == "" {
		log.Println("[WARNING] MISSING REDIS ENV: empty address. Change to default address localhost:6379")
		addr = "cpfd-redis:6379"
	}
	if pwd == "" {
		log.Println("[WARNING] MISSING REDIS ENV: empty password. Change to empty password")
		pwd = ""
	}
	if db == "" {
		log.Println("[WARNING] MISSING REDIS ENV: empty db. Change to default db '0'")
		db = "0"
	}
	intDb, err := strconv.Atoi(db)
	if err != nil {
		log.Println("[WARNING] INVALID Mysql: Change to default db '0'")
		intDb = 0
	}
	return conf{
		addr: addr,
		pwd:  pwd,
		db:   intDb,
	}
}

func (c conf) Info() {
	fmt.Println("========== REDIS ==========")
	fmt.Println("Address: ", c.addr)
	fmt.Println("Password: ", c.pwd)
	fmt.Println("Mysql: ", c.db)
}
