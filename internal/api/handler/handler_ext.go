package handler

import (
	"cpfd-back/internal/core"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLogsToCSV(ctx *gin.Context) {
	fPath := core.FileDir + "/data.csv"
	ctx.File(fPath)
}
