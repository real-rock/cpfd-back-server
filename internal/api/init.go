package api

import (
	"cpfd-back/internal/api/handler"
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/api/route"
	"cpfd-back/internal/api/service"
	"cpfd-back/internal/conf/grpc"
	"cpfd-back/internal/conf/mysqlDb"
	"cpfd-back/internal/conf/redisDb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	Engine *gin.Engine
	Mysql  *mysqlDb.DB
	Redis  *redisDb.Redis
}

func ConnMysql() *mysqlDb.DB {
	db := mysqlDb.New()
	//db.Migrate([]interface{}{&model.Activity{}, &model.Particle{}, &model.IndoorProperty{}})
	return db
}

func ConnRedis() *redisDb.Redis {
	return redisDb.New()
}

func NewRouter() *Router {
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
	if err := r.Engine.Run(":8080"); err != nil {
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
