package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/handlers"
	"github.com/mfs3curity/mynote/api/middlewares"
)

func Section(r *gin.RouterGroup) {
	section := handlers.NewSectionHandler()
	r.POST("/create", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), section.Create)
	r.GET("/name/:name", section.GetByName)
}
