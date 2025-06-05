package routes

import (
	"subscriptionmanager/internal/handlers"
	error "subscriptionmanager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(env *handlers.Env) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")
	router.Use(error.ErrorHandlerMiddleware())

	router.GET("/error", handlers.GetError)
	router.GET("/login", handlers.GetLogin)
	router.GET("/register", handlers.GetRegister)
	router.GET("/logout", handlers.GetLogout)
	router.GET("/", handlers.GetIndex)

	router.POST("/login", env.UserLogin)
	router.POST("/register", env.UserRegister)

	return router
}
