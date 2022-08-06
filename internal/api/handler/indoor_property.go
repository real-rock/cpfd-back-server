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

type IndoorPropertyHandler struct {
	Service *service.IndoorPropertyService
}

func NewIndoorPropertyHandler(s *service.IndoorPropertyService) *IndoorPropertyHandler {
	return &IndoorPropertyHandler{
		Service: s,
	}
}

func (h *IndoorPropertyHandler) CreateLog(ctx *gin.Context) {
	var ip model.IndoorProperty

	if err := ctx.ShouldBindJSON(&ip); err != nil {
		log.Printf("[ERROR] Failed to parse json: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	if err := h.Service.CreateLog(ip); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (h *IndoorPropertyHandler) GetLogToCSV(ctx *gin.Context) {
	var req struct {
		Start time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
		End   time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
	}

	if err := ctx.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := h.Service.GetLogToCSV(req.Start, req.End)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.File(filePath)
	defer func() {
		if err := os.Remove(filePath); err != nil {
			log.Println(err)
			return
		}
	}()
}
