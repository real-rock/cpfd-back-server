package handler

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/api/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateParticleLogs(ctx *gin.Context) {
	particle := service.CreateParticleLogParams{}

	if err := ctx.ShouldBindJSON(&particle); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateParticleLog(particle); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetParticleLogs(ctx *gin.Context) {
	req := repo.GetParticleLogsWithDatesParams{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(req)
	p, err := h.service.GetParticleLogsWithDates(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func (h *Handler) GetParticleLogsToCSV(ctx *gin.Context) {
	fPath, err := h.service.GetAllParticleLogsToCSV()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.File(fPath)
	defer func() {
		if err := os.Remove(fPath); err != nil {
			log.Println(err)
			return
		}
	}()
}

func (h *Handler) GetParticleLogsWithDates(ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, nil)
}

func (h *Handler) GetParticleLogsWithDatesToCSV(ctx *gin.Context) {
	req := repo.GetParticleLogsWithDatesParams{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fPath, err := h.service.GetParticleLogsWithDatesToCSV(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.File(fPath)
	defer func() {
		if err := os.Remove(fPath); err != nil {
			log.Println(err)
			return
		}
	}()
}

//func (h *ParticleHandler) GetLogToFile(ctx *gin.Context) {
//	var req struct {
//		Start   time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
//		End     time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
//		Method  string    `form:"method"`
//		Machine []string  `form:"machine"`
//	}
//	var filePath string
//	var err error

//	if err = ctx.ShouldBind(&req); err != nil {
//		log.Println(err.Error())
//		ctx.JSON(http.StatusBadRequest, err.Error())
//		return
//	}
//	if len(req.Machine) == 0 {
//		filePath, err = h.service.GetAllLogToFile(req.Start, req.End, req.Method)
//	} else {
//		filePath, err = h.service.GetLogToFile(req.Machine, req.Start, req.End)
//	}
//	ctx.File(filePath)
//	defer func() {
//		if err := os.Remove(filePath); err != nil {
//			log.Println(err)
//			return
//		}
//	}()
//}

//func (h *ParticleHandler) GetLogWithDates(ctx *gin.Context) {
//	var req struct {
//		Start time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
//		End   time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
//	}

//	if err := ctx.ShouldBind(&req); err != nil {
//		log.Println(err.Error())
//		ctx.JSON(http.StatusBadRequest, err.Error())
//		return
//	}
//	res, err := h.service.GetChartData(req.Start, req.End)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, err.Error())
//		return
//	}
//	ctx.JSON(http.StatusOK, res)
//}
