package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/handlers"
)

func User(r *gin.RouterGroup) {
	u := handlers.NewUserHandlers()
	r.POST("/login", u.Login)
}
