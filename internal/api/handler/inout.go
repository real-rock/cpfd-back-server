package handler

import (
	"cpfd-back/internal/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InoutHandler struct {
	Service *service.InoutService
}

func NewInoutHandler(s *service.InoutService) *InoutHandler {
	return &InoutHandler{
		Service: s,
	}
}

func (h *InoutHandler) GetInfo(ctx *gin.Context) {
	m, err := h.Service.GetCurrentInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, m)
}

func (h *InoutHandler) GetAllLogsToFile(ctx *gin.Context) {
	_, err := h.Service.GetAllLogsToFiles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.File("test.csv")
}

func (h *InoutHandler) GetLogs(ctx *gin.Context) {
	logs, err := h.Service.GetLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, logs)
}

func (h *InoutHandler) CreateLog(ctx *gin.Context) {
	var data struct {
		Name   string `json:"name"`
		Action bool   `json:"action"`
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Service.CreateLog(data.Name, data.Action); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully saved",
	})
}
