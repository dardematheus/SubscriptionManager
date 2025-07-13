package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"subscriptionmanager/internal/models"
	services "subscriptionmanager/internal/services"

	"github.com/gin-gonic/gin"
)

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

func (env *Env) GetIndex(c *gin.Context) {
	subscriptions, err := models.GetSubscriptions(c, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}
	costPerMonth, err := models.GetSumPerMonth(c, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}
	costPerYear, err := models.GetSumPerYear(c, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Subscriptions": subscriptions,
		"SumPerMonth":   costPerMonth,
		"SumPerYear":    costPerYear,
	})
}

func GetAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", nil)
}

func (env *Env) GetRemove(c *gin.Context) {
	subscriptions, err := models.GetSubscriptions(c, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}
	c.HTML(http.StatusOK, "remove.html", gin.H{
		"Subscriptions": subscriptions,
	})
}

func (env *Env) GetLogout(c *gin.Context) {
	cookie, err := c.Cookie("session_cookie")
	if err != nil {
		c.Error(errors.New("user not logged in")).SetMeta(400)
		http.Redirect(c.Writer, c.Request, "/unauthorized", http.StatusSeeOther)
		return
	}
	services.DeleteSession(c, env.DB, cookie)
	c.HTML(http.StatusOK, "login.html", nil)
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
	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
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
	http.Redirect(c.Writer, c.Request, "/login", http.StatusSeeOther)
	return
}

func (env *Env) AddSubscription(c *gin.Context) {
	subscription := c.PostForm("name")
	date := c.PostForm("date")
	costStr := c.PostForm("cost")
	cost, err := strconv.ParseFloat(costStr, 64)
	if err != nil {
		c.Error(errors.New("NULL Cost")).SetMeta(400)
		return
	}

	if subscription == "" || cost <= 0 || date == "" {
		c.Error(errors.New("all Fields Required")).SetMeta(400)
		return
	}
	if err = models.AddSubscription(subscription, date, cost, c, env.DB); err != nil {
		c.Error(err).SetMeta(400)
	}
	http.Redirect(c.Writer, c.Request, "/", http.StatusSeeOther)
	return
}

func (env *Env) RemoveSubscription(c *gin.Context) {
	subscription := c.PostForm("subscription")

	if subscription == "" {
		c.Error(errors.New("subscription is required")).SetMeta(400)
		return
	}
	if err := models.RemoveSubscription(subscription, c, env.DB); err != nil {
		c.Error(err).SetMeta(400)
	}
}
