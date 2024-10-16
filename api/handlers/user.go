package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/services"
)

type UserHandlers struct {
	service *services.UserService
}

func NewUserHandlers() *UserHandlers {
	return &UserHandlers{
		service: services.NewUserService(),
	}
}

func (h *UserHandlers) Login(ctx *gin.Context) {
	req := new(dto.Login)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	t, err := h.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, t)
}
