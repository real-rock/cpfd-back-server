package handler

import (
	"cpfd-back/internal/api/repo"
	"cpfd-back/internal/api/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateIndoorPropertyLog(ctx *gin.Context) {
	req := service.CreateIndoorPropertyLogParams{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	if err := h.service.CreateIndoorPropertyLog(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetIndoorPropertyLogs(ctx *gin.Context) {
	req := repo.GetIndoorPropertyLogsParams{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ip, err := h.service.GetIndoorPropertyLogs(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, ip)
}

func (h *Handler) GetIndoorPropertyLogsToCSV(ctx *gin.Context) {
	req := repo.GetIndoorPropertyLogsParams{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fPath, err := h.service.GetIndoorPropertyLogsToCSV(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.File(fPath)
	defer func() {
		if err := os.Remove(fPath); err != nil {
			return
		}
	}()
}
