package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLogsToCSV(ctx *gin.Context) {
	fPath, err := h.service.GetAllLogsToCSV()
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
