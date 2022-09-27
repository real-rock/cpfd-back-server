package route

import (
	"cpfd-back/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func Set(router *gin.RouterGroup, h handler.HandlerManager) {
	router.GET("/machines", h.GetMachine)
	router.POST("/machines", h.CreateMachine)

	router.GET("/logs/file", h.GetAllLogsToCSV)

	router.GET("/logs/activity", h.GetActivityLogs)
	router.GET("/logs/file/activity", h.GetActivityLogsToCSV)
	router.POST("/logs/activity", h.CreateActivityLog)
	router.GET("/logs/activity/state", h.GetCurrentActivity)

	router.GET("/logs/particle", h.GetParticleLogs)
	router.GET("/logs/file/particle", h.GetParticleLogsToCSV)
	router.POST("/logs/particle", h.CreateParticleLogs)

	router.POST("/logs/indoor-property", h.CreateIndoorPropertyLog)
	router.GET("/logs/indoor-property", h.GetIndoorPropertyLogs)
	router.GET("/logs/file/indoor-property", h.GetIndoorPropertyLogsToCSV)
}
