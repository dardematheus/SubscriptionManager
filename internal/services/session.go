package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"io"

	"github.com/gin-gonic/gin"
)

func generateSessionID() (string, error) {
	const size = 16
	bytes := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateSession(c *gin.Context, userID int, db *sql.DB) error {
	cookie, err := c.Cookie("session_cookie")

	if err == nil {
		_, err = db.Exec("DELETE FROM sessions WHERE id = ?", cookie)
		c.SetCookie("session_cookie", "", -1, "/", "localhost", false, true)
	}

	cookie, err = generateSessionID()
	_, err = db.Exec("INSERT INTO sessions (id, user_id) VALUES (? , ? )", cookie, userID)

	if err == nil {
		c.SetCookie("session_cookie", cookie, 86400, "/", "localhost", false, true)
		return nil
	}
	return err
}

func ValidateSession(db *sql.DB, cookie string) bool {
	_, err := db.Exec("SELECT user_id FROM sessions WHERE id = ?", cookie)
	return err == nil
}
