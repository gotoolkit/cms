package middlewares

import (
	"net/http"
	"time"

	"github.com/gotoolkit/cms/pkg/model"

	"github.com/gotoolkit/cms/pkg/database"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// DefaultJwt ...
var DefaultJwt = &jwt.GinJWTMiddleware{
	Realm:      "sme Realm",
	Key:        []byte("sme2018app"),
	Timeout:    time.Minute,
	MaxRefresh: time.Hour,
	Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
		total, err := database.GetDB().Where("username = ?", userId).Where("password = ?", password).Count(&model.PurchaseUser{})
		if err != nil {
			return userId, false
		}
		if total != 1 {
			return userId, false
		}
		return userId, true
	},
	Authorizator: func(userId string, c *gin.Context) bool {
		total, err := database.GetDB().Where("username = ?", userId).Count(&model.PurchaseUser{})
		if err != nil {
			return false
		}
		if total != 1 {
			return false
		}
		return true
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": message,
		})
	},
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

// Jwt ...
func Jwt() *jwt.GinJWTMiddleware {
	return DefaultJwt
}
