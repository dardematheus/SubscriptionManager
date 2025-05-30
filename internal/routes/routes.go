package routes

import (
	"database/sql"
	error "subscriptionmanager/internal/error"
	"subscriptionmanager/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Initializes GinRouter with middleware + handlers
func InitRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")
	router.Use(error.ErrorHandler())

	env := handlers.NewEnv(db)

	//GET Requests
	router.GET("/login", handlers.GetLogin)
	router.GET("/register", handlers.GetRegister)
	router.GET("/", handlers.GetIndex)
	router.GET("/edit", handlers.GetEdit)
	router.GET("/logout", handlers.GetLogout)

	//POST Requests
	router.POST("/login", env.UserLogin)
	router.POST("/register", env.UserRegister)

	return router
}
