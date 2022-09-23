package handler

import (
	"cpfd-back/internal/api/service"

	"github.com/gin-gonic/gin"
)

type HandlerManager interface {
	CreateMachine(ctx *gin.Context)
	GetMachine(ctx *gin.Context)

	GetAllLogsToCSV(ctx *gin.Context)

	CreateActivityLog(ctx *gin.Context)
	GetCurrentActivity(ctx *gin.Context)
	GetActivityLogs(ctx *gin.Context)
	GetActivityLogsToCSV(ctx *gin.Context)

	CreateParticleLogs(ctx *gin.Context)
	GetParticleLogs(ctx *gin.Context)
	GetParticleLogsToCSV(ctx *gin.Context)
	GetParticleLogsWithDates(ctx *gin.Context)
	GetParticleLogsWithDatesToCSV(ctx *gin.Context)

	CreateIndoorPropertyLog(ctx *gin.Context)
	GetIndoorPropertyLogs(ctx *gin.Context)
	GetIndoorPropertyLogsToCSV(ctx *gin.Context)
}

func New(s service.ServiceManager) HandlerManager {
	return &Handler{
		service: s,
	}
}

type Handler struct {
	service service.ServiceManager
}
