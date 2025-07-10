package routes

import (
	"subscriptionmanager/internal/handlers"
	middleware "subscriptionmanager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(env *handlers.Env) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")
	router.Use(middleware.ErrorHandlerMiddleware())

	authRouter := router.Group("/")
	authRouter.Use(middleware.AuthMiddleware(env.DB))
	{
		authRouter.GET("/", handlers.GetIndex)
		authRouter.GET("/add", handlers.GetAdd)
		authRouter.GET("/remove", handlers.GetRemove)
		authRouter.GET("logout", handlers.GetLogout)
		authRouter.POST("/add", env.AddSubscription)
		authRouter.POST("/remove", env.RemoveSubscription)
	}

	router.GET("/login", handlers.GetLogin)
	router.GET("/unauthorized", handlers.GetError)
	router.GET("/register", handlers.GetRegister)

	router.POST("/login", env.UserLogin)
	router.POST("/register", env.UserRegister)

	return router
}
