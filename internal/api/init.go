package api

import (
	"cpfd-back/internal/api/handler"
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/api/route"
	"cpfd-back/internal/api/service"
	"cpfd-back/internal/conf/grpc"
	"cpfd-back/internal/conf/mysqlDb"
	"cpfd-back/internal/conf/redisDb"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Router struct {
	Engine *gin.Engine
	Mysql  *mysqlDb.DB
	Redis  *redisDb.Redis
}

func ConnMysql() *mysqlDb.DB {
	db := mysqlDb.New()
	return db
}

func ConnRedis() *redisDb.Redis {
	return redisDb.New()
}

func NewRouter() *Router {
	if viper.Get("app.production") == true {
		gin.SetMode(gin.ReleaseMode)
		ginLog, err := os.Create("internal/log/gin.log")
		if err == nil {
			gin.DefaultWriter = io.MultiWriter(ginLog)
		} else {
			log.Println("[ERROR] failed to create gin log file: ", err.Error())
		}
	}
	router := gin.Default()
	router.Use(cors.Default())
	r := &Router{
		Engine: router,
		Mysql:  ConnMysql(),
		Redis:  ConnRedis(),
	}
	r.Set()
	return r
}

func (r *Router) Run() {
	port := viper.GetString("app.port")
	if err := r.Engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("[ERROR] error while running server: %v", err)
	}
}

func (r *Router) Set() {
	r.setInout()
	r.setParticle()
	r.setIndoorProperty()
}

func (r *Router) setInout() {
	re := repo.NewInoutRepo(r.Mysql.DB, r.Redis.DB)
	s := service.NewInoutService(re)
	h := handler.NewInoutHandler(s)
	route.SetInout(r.Engine.Group("/v1"), h)
}

func (r *Router) setParticle() {
	re := repo.NewParticleRepo(r.Mysql.DB)
	s := service.NewParticleService(re, grpc.New())
	h := handler.NewParticleHandler(s)
	route.SetParticle(r.Engine.Group("/v1"), h)
}

func (r *Router) setIndoorProperty() {
	re := repo.NewIndoorPropertyRepo(r.Mysql.DB)
	s := service.NewIndoorPropertyService(re)
	h := handler.NewIndoorPropertyHandler(s)
	route.SetIndoorProperty(r.Engine.Group("/v1"), h)
}
