package routes

import (
	"database/sql"
	error "subscriptionmanager/internal/error"
	"subscriptionmanager/internal/handlers"

	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
)

var session *scs.Session

func InitRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")
	router.Use(error.ErrorHandler())

	env := handlers.NewEnv(db)

	authRouter := router.Group("/auth")
	{
		authRouter.Use()
		authRouter.GET("/")
		authRouter.GET("edit")
	}

	router.GET("/login", handlers.GetLogin)
	router.GET("/register", handlers.GetRegister)
	router.GET("/logout", handlers.GetLogout)

	router.POST("/login", env.UserLogin)
	router.POST("/register", env.UserRegister)

	return router
}
