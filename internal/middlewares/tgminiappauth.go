package middlewares

import (
	"fmt"
	"net/http"
	"resedist/pkg/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func TgAuthMiddleware() gin.HandlerFunc {

	botToken := config.Get().Telegram.BotToken
	expIn := 24 * time.Hour

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

		// پارس initData
		initData, err := initdata.Parse(initDataStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid initData format"})
			c.Abort()
			return
		}

		fmt.Println(initData)

		//initdata.Validate(initDataStr, botToken, expIn)
		// اعتبارسنجی
		if err := initdata.Validate(initDataStr, botToken, expIn); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid initData: " + err.Error()})
			c.Abort()
			return
		}

		// اگر معتبر باشه، اطلاعات کاربر رو به context اضافه کنید
		c.Set("user_id", initData.User.ID)
		c.Set("username", initData.User.Username)

		//user_id, _ := c.Get("user_id")
		//username, _ := c.Get("user_id")
		//
		//fmt.Println("Authenticated user_id:", user_id)
		//fmt.Println("Authenticated username:", username)
		c.Next()
	}
}
