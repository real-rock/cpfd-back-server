package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateActivityLog(ctx *gin.Context) {
	var data struct {
		Name   string `json:"name"`
		Action bool   `json:"action"`
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateActivityLog(data.Name, data.Action); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) GetCurrentActivity(ctx *gin.Context) {
	m, err := h.service.GetCurrentActivity()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, m)
}

func (h *Handler) GetActivityLogs(ctx *gin.Context) {
	a, err := h.service.GetActivityLogs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, a)
}

func (h *Handler) GetActivityLogsToCSV(ctx *gin.Context) {
	fn, err := h.service.GetActivityLogsToCSV()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer os.Remove(fn)
	ctx.File(fn)
}
