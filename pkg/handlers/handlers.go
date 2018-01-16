package handlers

import (
	"log"
	"net/http"
	"net/http/pprof"
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

	r.GET("/version", gin.WrapF(home(buildTime, commit, release)))
	r.GET("/healthz", gin.WrapF(healthz))
	r.GET("/readyz", gin.WrapF(readyz(isReady)))

	if gin.IsDebugging() {
		debugRouter(r)
	}

	r.POST("/login", middlewares.Jwt().LoginHandler)
	r.POST("/signup", register)
	r.GET("message/:phone", sendMessage)

	auth := r.Group("/auth")
	auth.Use(middlewares.Jwt().MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", middlewares.Jwt().RefreshHandler)

		companyRest(auth, &company{})
	}

	return r
}

func companyRest(rg *gin.RouterGroup, c *company) {

	rg.GET("/company", c.GetAll())
	rg.GET("/company/:id", c.Get())
	rg.POST("/company/:id", c.Create())
	rg.DELETE("/company/:id", c.Delete())
}

func debugRouter(r *gin.Engine) {
	r.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/heap", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/block", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/mutex", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	r.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	r.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	r.POST("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	r.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

func abortWithStatus(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}
