package routes

import (
	error "SubscriptionManager/internal/error"
	"SubscriptionManager/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Initializes GinRouter with middleware + handlers
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")
	router.Use(error.ErrorHandler())

	//GET Requests
	router.GET("/login", handlers.GetLogin)
	router.GET("/register", handlers.GetRegister)
	router.GET("/", handlers.GetIndex)
	router.GET("/edit", handlers.GetEdit)
	router.GET("/logout", handlers.GetLogout)

	//POST Requests
	router.POST("/login", handlers.UserLogin)
	router.POST("/register", handlers.UserRegister)

	return router
}
