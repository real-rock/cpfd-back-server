package main

import (
	"fmt"
	"log"
	"os"

	logger "cpfd-back/internal/log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type app struct {
	Name       string
	Version    string
	Production bool
	Storage    []string
	Port       string
	Logger     string
}

func NewApp() *app {
	_ = os.Setenv("TZ", "Asia/Seoul")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("[ERROR] failed to read config file: ", err.Error())
	}
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalln("[ERROR] failed to load env file: ", err.Error())
	}
	fmt.Println(os.Getwd())
	logger.SetLogger()
	return &app{
		Name:       viper.GetString("app.name"),
		Version:    viper.GetString("app.version"),
		Production: viper.GetBool("app.production"),
		Storage:    viper.GetStringSlice("app.storage"),
		Port:       viper.GetString("app.port"),
		Logger:     viper.GetString("app.logger"),
	}
}

func (a *app) PrintAppInfo() {
	fmt.Println("======= INFO =======")
	fmt.Println("App name: ", viper.GetString("app.name"))
	fmt.Println("App version: ", viper.GetString("app.version"))
	fmt.Println("App production: ", viper.GetBool("app.production"))
	fmt.Println("App storages: ", viper.GetStringSlice("app.storage"))
	fmt.Println("App port: ", viper.GetString("app.port"))
	fmt.Println("App logger: ", viper.GetString("app.logger"))
}
