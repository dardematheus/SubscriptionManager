package middleware

import (
	"database/sql"
	"net/http"
	"subscriptionmanager/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_cookie")
		if err != nil {
			http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
			return
		}
		if !services.ValidateSession(db, cookie) {
			http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
			return
		}
		c.Next()
	}
}
