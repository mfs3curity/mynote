package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/middlewares"
	"github.com/mfs3curity/mynote/api/routers"
)

func InitServer() {
	r := gin.New()
	r.Static("/images/uploads", "./images/uploads")
	r.Use(middlewares.Cors(), middlewares.CodedHeader())
	// router
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		routers.User(auth)
	}
	{
		post := api.Group("/post")
		routers.Post(post)
	}
	{
		section := api.Group("/section")
		routers.Section(section)
	}
	r.Run(":1337")
}
