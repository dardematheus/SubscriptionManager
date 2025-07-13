package models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddSubscription(subscription, date string, cost int, c *gin.Context, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO subscriptions (name, date, cost, user_id) VALUES (?, ?, ?, ?)",
		subscription, date, cost, c.GetInt("UserID"))
	if err != nil {
		return err
	}
	return nil
}

func RemoveSubscription(subscription string, c *gin.Context, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM subscriptions WHERE name = ? AND user_id = ?", subscription, c.GetInt("UserID"))
	if err != nil {
		return err
	}
	return nil
}

func GetSubscriptions(c *gin.Context, db *sql.DB) {

}
