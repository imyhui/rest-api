package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/handler/sd"
	"rest-api/handler/user"
	"rest-api/router/middleware"
)

// Load loads the middleeares, routes, handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares
	// Set Header
	// Clear up Panic
	g.Use(gin.Recovery())
	// Browser does not cache
	g.Use(middleware.NoCache)
	// Cross-domain Options request
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
	return g
}
