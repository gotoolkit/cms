package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/models"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

func register(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Missing Username or Password",
		})
		return
	}

	if !database.GetDB().Where("username = ?", user.Username).First(&models.User{}).RecordNotFound() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "User already registed",
		})
		return
	}

	database.GetDB().Create(&user)

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)

	expire := time.Now().Add(time.Hour)
	claims["id"] = user.Username
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte("secret key"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Create JWT Token faild",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}
