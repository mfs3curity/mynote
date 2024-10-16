package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/handlers"
	"github.com/mfs3curity/mynote/api/middlewares"
)

func Post(r *gin.RouterGroup) {
	post := handlers.NewPostHandlers()
	r.POST("/create", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), post.Create)
	r.GET("/id/:id", post.GetByID)
	r.POST("/upload", middlewares.Authentication(), middlewares.Authorization([]string{"admin"}), post.Upload)
}
