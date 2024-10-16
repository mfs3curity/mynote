package middlewares

import "github.com/gin-gonic/gin"

func CodedHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("development", "MF+0xle0 -> Telegram:https://t.me/gofirebot + https://t.me/xLe0x")
		c.Header("version", "1.0 Beta")
		c.Next()
	}
}
