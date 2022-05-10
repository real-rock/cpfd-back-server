package route

import (
	"cpfd-back/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func SetInout(router *gin.RouterGroup, handler *handler.InoutHandler) {
	router.GET("/logs/log", handler.GetLogs)
	router.GET("/logs/file", handler.GetAllLogsToFile)
	router.POST("/logs/log", handler.CreateLog)
	router.GET("/logs/info", handler.GetInfo)
}

func SetParticle(router *gin.RouterGroup, handler *handler.ParticleHandler) {
	router.GET("/logs/particle", handler.GetLogs)
	router.GET("/logs/chart/particle", handler.GetLogWithDates)
	router.POST("/logs/particle", handler.CreateLog)
}
