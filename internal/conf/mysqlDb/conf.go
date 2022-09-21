package mysqlDb

import (
	"fmt"
	"os"
)

type conf struct {
	dial string
	user string
	pwd  string
	host string
	port string
	name string
}

func newConf() conf {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USR")
	pwd := os.Getenv("MYSQL_PWD")
	name := os.Getenv("MYSQL_DB")

	return conf{
		dial: "mysqlDb",
		user: user,
		pwd:  pwd,
		host: host,
		port: port,
		name: name,
	}
}

func (c conf) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.user, c.pwd, c.host, c.port, c.name)
}

func (c conf) Info() {
	fmt.Println("========== Mysql ==========")
	fmt.Println("Dial: ", c.dial)
	fmt.Println("User: ", c.user)
	fmt.Println("Password: ", c.pwd)
	fmt.Println("Host: ", c.host)
	fmt.Println("Port: ", c.port)
	fmt.Println("Name: ", c.name)
}
