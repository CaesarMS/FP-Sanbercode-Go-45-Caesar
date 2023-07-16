package middleware

import (
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, config.Response{
				Error: "invalid access token",
			})
			c.Abort()
			return
		}

		if len(strings.Split(token, " ")) == 2 {
			token = strings.Split(token, " ")[1]
		}

		isTokenNotValid := util.VerifyJWT(token, "user")

		if isTokenNotValid != nil {
			c.JSON(http.StatusUnauthorized, config.Response{
				Error: isTokenNotValid.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
