package middleware

import (
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"github.com/gin-gonic/gin"
)

func AdminKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminKey, err := util.GetEnv("API_KEY_ADMIN")
		if err != nil {
			c.JSON(http.StatusUnauthorized, config.Response{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if adminKey != c.Request.Header.Get("x-api-key") {
			c.JSON(http.StatusUnauthorized, config.Response{
				Error: "invalid api key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminTokenMiddleware() gin.HandlerFunc {
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

		isTokenNotValid := util.VerifyJWT(token, "admin")

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
