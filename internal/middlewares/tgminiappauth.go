package middlewares

import (
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"resedist/pkg/config"
	"strings"
)

func TgAuthMiddleware() gin.HandlerFunc {

	botToken := config.Get().Telegram.BotToken
	expIn := config.Get().Telegram.TokenExpr

	return func(c *gin.Context) {
		initDataStr := c.GetHeader("Authorization")
		if strings.HasPrefix(initDataStr, "Bearer ") {
			initDataStr = strings.TrimPrefix(initDataStr, "Bearer ")
		}
		if initDataStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No initData provided"})
			c.Abort()
			return
		}

		initData, err := initdata.Parse(initDataStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid initData format"})
			c.Abort()
			return
		}

		if err := initdata.Validate(initDataStr, botToken, expIn); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid initData: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("tg_user", initData)
		c.Next()
	}
}
