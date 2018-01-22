package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cms/pkg/database"
	"github.com/gotoolkit/cms/pkg/model"
	"github.com/gotoolkit/cms/pkg/teleport"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

func register(c *gin.Context) {
	var user model.PurchaseUser

	err := c.BindJSON(&user)
	if err != nil {
		log.Println(err)
		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Missing Username or Password: %v", err))
		return
	}

	has, err := database.GetDB().Where("username = ?", user.Username).Get(&model.PurchaseUser{})
	if err != nil {
		log.Println(err)

		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Database connection failed: %v", err))
		return
	}
	if has {
		log.Println(err)

		abortWithStatus(c, http.StatusBadRequest, "User has already registed")
		return
	}
	user.UsernameCanonical = user.Username
	user.Email= user.Username
	user.EmailCanonical = user.Email
	user.ConfirmationToken = strconv.Itoa(int(time.Now().UnixNano()))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Roles = "a:0:{}"
	user.Configs = "a:0:{}"
	_, err = database.GetDB().Insert(&user)

	if err != nil {
		log.Println(err)
		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Create new user failed: %v", err))
		return
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)

	expire := time.Now().Add(time.Hour)
	claims["id"] = user.Username
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte("sme360 secret key"))
	if err != nil {
		log.Println(err)
		abortWithStatus(c, http.StatusUnauthorized, fmt.Sprintf("Create JWT Token faild: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

func sendMessage(c *gin.Context) {
	phone := c.Param("phone")
	has, err := database.GetDB().Where("mobile = ?", phone).Get(&model.PurchaseUser{})
	if err != nil {
		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Database connection failed: %v", err))
		return
	}

	if has {
		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Phone already registed:"))
		return
	}

	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999-99999) + 99999
	err = teleport.SendMessage(fmt.Sprintf("phone: %s \ncode: %d", phone, code))

	if err != nil {
		abortWithStatus(c, http.StatusBadRequest, fmt.Sprintf("Send telegram message failed: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": code,
	})
}
