package models

import (
	"database/sql"
	"errors"
	"log"

	"subscriptionmanager/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Subscription struct {
	Name string
	Date string
	Cost float64
}

func AddSubscription(subscription, date string, cost float64, c *gin.Context, db *sql.DB) error {
	value := int(cost * 100)

	sessionAny, exists := c.Get("session")
	if !exists {
		return errors.New("session not found in context")
	}
	session, ok := sessionAny.(middleware.SessionInfo)
	if !ok {
		return errors.New("invalid session type")
	}
	userID := session.UserID

	_, err := db.Exec("INSERT INTO subscriptions (name, date, cost, user_id) VALUES (?, ?, ?, ?)",
		subscription, date, value, userID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveSubscription(subscription string, c *gin.Context, db *sql.DB) error {
	sessionAny, exists := c.Get("session")
	if !exists {
		return errors.New("session not found in context")
	}
	session, ok := sessionAny.(middleware.SessionInfo)
	if !ok {
		return errors.New("invalid session type")
	}
	userID := session.UserID

	_, err := db.Exec("DELETE FROM subscriptions WHERE name = ? AND user_id = ?", subscription, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetSubscriptions(c *gin.Context, db *sql.DB) ([]Subscription, error) {
	sessionAny, exists := c.Get("session")
	if !exists {
		return nil, errors.New("session not found in context")
	}
	session, ok := sessionAny.(middleware.SessionInfo)
	if !ok {
		return nil, errors.New("invalid session type")
	}
	userID := session.UserID

	rows, err := db.Query("SELECT name, date, cost FROM subscriptions WHERE user_id = ?", userID)
	if err != nil {
		c.Error(err).SetMeta(400)
		return nil, err
	}
	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		var name, date string
		var cost int
		err = rows.Scan(&name, &date, &cost)
		if err != nil {
			c.Error(err).SetMeta(400)
			return nil, err
		}
		sub := Subscription{
			Name: name,
			Date: date,
			Cost: float64(cost) / 100,
		}
		subscriptions = append(subscriptions, sub)
		log.Println(subscriptions)
	}
	return subscriptions, nil
}

func GetSumPerMonth(c *gin.Context, db *sql.DB) (float64, error) {
	sessionAny, exists := c.Get("session")
	if !exists {
		return 0, errors.New("session not found in context")
	}
	session, ok := sessionAny.(middleware.SessionInfo)
	if !ok {
		return 0, errors.New("invalid session type")
	}
	userID := session.UserID

	var total sql.NullInt64
	err := db.QueryRow("SELECT SUM(cost) FROM subscriptions WHERE user_id = ?", userID).Scan(&total)
	if err != nil {
		return 0, err
	}
	if !total.Valid {
		return 0, nil
	}
	return float64(total.Int64) / 100, nil
}

func GetSumPerYear(c *gin.Context, db *sql.DB) (float64, error) {
	sessionAny, exists := c.Get("session")
	if !exists {
		return 0, errors.New("session not found in context")
	}
	session, ok := sessionAny.(middleware.SessionInfo)
	if !ok {
		return 0, errors.New("invalid session type")
	}
	userID := session.UserID

	var total sql.NullInt64
	err := db.QueryRow("SELECT SUM(cost) FROM subscriptions WHERE user_id = ?", userID).Scan(&total)
	if err != nil {
		return 0, err
	}
	if !total.Valid {
		return 0, nil
	}
	return float64(total.Int64) / 100 * 12, nil
}
