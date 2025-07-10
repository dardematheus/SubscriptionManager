package handlers

import (
	"errors"
	"net/http"
	"subscriptionmanager/internal/models"
	services "subscriptionmanager/internal/services"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string
	Password string
}

type SubscriptionInfo struct {
	Name string
	Cost int
	Date string
}

// GET Handlers
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetError(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", nil)
}

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", nil)
}

func GetRemove(c *gin.Context) {
	c.HTML(http.StatusOK, "remove.html", nil)
}

func GetLogout(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/login.html", nil)
}

func GetUnauthorized(c *gin.Context) {
	c.HTML(http.StatusUnauthorized, "error.html", nil)
}

// POST Handlers
func (env *Env) UserLogout(c *gin.Context) {
	cookie, err := c.Cookie("session_cookie")
	if err != nil {
		c.Error(errors.New("user not logged in")).SetMeta(400)
		http.Redirect(c.Writer, c.Request, "/unauthrorized", http.StatusSeeOther)
	}
	services.DeleteSession(c, env.DB, cookie)
	http.Redirect(c.Writer, c.Request, "/login", http.StatusSeeOther)
}

func (env *Env) UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.Error(errors.New("all fields are required")).SetMeta(400)
		return
	}
	userID, err := models.LoginUser(username, password, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}

	if err = services.CreateSession(c, userID, env.DB); err != nil {
		c.Error(err).SetMeta(400)
		return
	}
	http.Redirect(c.Writer, c.Request, "/index", http.StatusSeeOther)
}

func (env *Env) UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	if username == "" || password == "" || confirmPassword == "" {
		c.Error(errors.New("all fields are required")).SetMeta(400)
		return
	}
	if password != confirmPassword {
		c.Error(errors.New("passwords do not match")).SetMeta(400)
		return
	}
	if err := models.RegisterUser(username, password, env.DB); err != nil {
		c.Error(err).SetMeta(400)
	}
}

func (env *Env) AddSubscription(c *gin.Context) {
	subscription := c.PostForm("subscription")
	cost := c.PostForm("cost")
	date := c.PostForm("Date")

	if subscription == "" || cost == "" || date == "" {
		c.Error(errors.New("All Fields Required")).SetMeta(400)
		return
	}
	if err := models.AddSubscription(subscription, cost, date, env.DB); err != nil {
		c.Error(err).SetMeta(400)
	}
}

func (env *Env) RemoveSubscription(c *gin.Context) {
	subscription := c.PostForm("subscription")

	if subscription == "" {
		c.Error(errors.New("Subscription is required")).SetMeta(400)
		return
	}
	if err := models.RemoveSubscription(subscription, env.DB); err != nil {
		c.Error(err).SetMeta(400)
	}
}
