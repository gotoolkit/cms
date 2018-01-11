package middlewares

import (
	"time"

	"github.com/gotoolkit/cms/pkg/models"

	"github.com/gotoolkit/cms/pkg/database"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Jwt() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if !database.GetDB().Where("username = ?", userId).Where("password = ?", password).Find(&models.User{}).RecordNotFound() {
				return userId, true
			}
			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return !database.GetDB().Where("username = ?", userId).First(&models.User{}).RecordNotFound()
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
