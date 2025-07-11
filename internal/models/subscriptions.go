package models

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddSubscription(subscription, date string, cost int, c *gin.Context, db *sql.DB) error {

}

func RemoveSubscription(subscription string, c *gin.Context, db *sql.DB) error {

}
