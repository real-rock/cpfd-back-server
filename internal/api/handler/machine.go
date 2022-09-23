package handler

import (
	"cpfd-back/internal/api/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMachine(ctx *gin.Context) {
	machine := repo.CreateMachineParams{}

	if err := ctx.ShouldBindJSON(&machine); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.CreateMachine(machine); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetMachine(ctx *gin.Context) {
	var req struct {
		Id  string `json:"id"`
		Num int    `json:"num"`
	}
	m, err := h.service.GetMachine(req.Id, req.Num)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, m)
}
