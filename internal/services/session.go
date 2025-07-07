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

	//Sets the Cookie if doesnt exist
	if err != nil {
		cookie, err = generateSessionID()
		_, err = db.Exec("INSERT INTO sessions (id, user_id) VALUES (? , ? )", cookie, userID)
		c.SetCookie("session_cookie", cookie, 86400, "/", "localhost", false, true)
		return nil
	}
	return nil
}
