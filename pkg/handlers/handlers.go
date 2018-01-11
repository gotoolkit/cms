package handlers

import (
	"log"
	"sync/atomic"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/gotoolkit/cms/pkg/middlewares"
)

func Router(buildTime, commit, release string) *gin.Engine {

	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(5 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()
	r := gin.Default()

	r.GET("/home", gin.WrapF(home(buildTime, commit, release)))
	r.GET("/healthz", gin.WrapF(healthz))
	r.GET("/readyz", gin.WrapF(readyz(isReady)))

	r.POST("/login", middlewares.Jwt().LoginHandler)
	r.POST("/signup", register)

	auth := r.Group("/auth")
	auth.Use(middlewares.Jwt().MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", middlewares.Jwt().RefreshHandler)
	}

	return r
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}
