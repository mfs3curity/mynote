package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mfs3curity/mynote/services"
)

func Authentication() gin.HandlerFunc {
	var tokenService = services.NewTokenService()

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader("Authorization")
		token := strings.Split(auth, " ")
		if auth == "" {
			err = errors.New("token required")
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = errors.New("token expired")
				default:
					err = errors.New("invalid token")
				}
			}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
				"msg": "StatusUnauthorized",
			})
			return
		}

		c.Set("UserId", claimMap["UserId"])
		c.Set("username", claimMap["username"])
		c.Set("roles", claimMap["roles"])
		c.Set("exp", claimMap["exp"])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "forbidden Authorization"})
			return
		}
		rolesVal := c.Keys["roles"]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "forbidden Authorization no permission"})
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "forbidden Authorization"})
	}
}
