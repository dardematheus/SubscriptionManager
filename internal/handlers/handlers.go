package handlers

import (
	"errors"
	"net/http"
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

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/edit.html", nil)
}

func GetLogout(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/login.html", nil)
}

// POST Handlers
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
	err := models.RegisterUser(username, password, env.DB)
	if err != nil {
		c.Error(err).SetMeta(400)
		return
	}
}
