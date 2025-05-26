package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET Handlers
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/index.html", nil)
}

func GetEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/edit.html", nil)
}

func GetLogout(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/login.html", nil)
}

// POST Handlers
func UserLogin(c *gin.Context) {

}

func UserRegister(c *gin.Context) {
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
}
