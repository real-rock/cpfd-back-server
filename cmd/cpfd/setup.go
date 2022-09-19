package main

import (
	"fmt"
	"log"

	logger "cpfd-back/internal/log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type app struct {
	Name       string
	Version    string
	Production bool
	Storage    []string
	Logger     string
}

func printAppInfo() {
	fmt.Println("======= INFO =======")
	fmt.Println("App name: ", viper.Get("app.name"))
	fmt.Println("App version: ", viper.Get("app.version"))
	fmt.Println("App production: ", viper.Get("app.production"))
	fmt.Println("App storages: ", viper.Get("app.stroage"))
	fmt.Println("App port: ", viper.Get("app.port"))
}

func setApp() {
	viper.SetConfigFile("app.yaml")
	if err := godotenv.Load("db.env"); err != nil {
		log.Fatalln("[ERROR] Failed to load env file: ", err.Error())
	}
	printAppInfo()
	logger.SetLogger()
}
