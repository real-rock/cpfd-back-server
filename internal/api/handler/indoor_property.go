package handler

import (
	"cpfd-back/internal/api/service"
	"cpfd-back/internal/core/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
