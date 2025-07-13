package middleware

import (
	"database/sql"
	"net/http"
	"subscriptionmanager/internal/services"

	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserID    int
	SessionID string
}

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID int
		cookie, err := c.Cookie("session_cookie")
		if err != nil {
			http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
			return
		}
		if !services.ValidateSession(db, cookie) {
			http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
			return
		}

		err = db.QueryRow("SELECT user_id FROM sessions WHERE id = ?", cookie).Scan(&userID)
		if err != nil {
			http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
		}

		session := SessionInfo{
			UserID:    userID,
			SessionID: cookie,
		}

		c.Set("session", session)
		c.Next()
	}
}
