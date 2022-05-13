package handler

import (
	"cpfd-back/internal/api/service"
	"cpfd-back/internal/core/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type ParticleHandler struct {
	service *service.ParticleService
}

func NewParticleHandler(s *service.ParticleService) *ParticleHandler {
	return &ParticleHandler{
		service: s,
	}
}

func (h *ParticleHandler) GetLogs(ctx *gin.Context) {
	particles, err := h.service.GetLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, particles)
}

func (h *ParticleHandler) GetLogToFile(ctx *gin.Context) {
	var req struct {
		Start time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
		End   time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := h.service.GetLogToFile(req.Start, req.End)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//filePath := core.FileDir + "/" + fileName + ".csv"
	ctx.File(filePath)
	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Println(err)
			return
		}
	}()
}

func (h *ParticleHandler) GetLogWithDates(ctx *gin.Context) {
	var req struct {
		Start time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
		End   time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.service.GetChartData(req.Start, req.End)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *ParticleHandler) CreateLog(ctx *gin.Context) {
	particle := model.Particle{}

	if err := ctx.ShouldBindJSON(&particle); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateLog(particle); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
