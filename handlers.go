package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Remember bool   `form:"remember"`
}

func (a *App) handleIndex(c *gin.Context) {
	_, err := c.Cookie("username")
	if err != nil {
		c.HTML(http.StatusOK, "login.go.html", gin.H{})
		return
	}
	c.HTML(http.StatusOK, "index.go.html", gin.H{"message": "HellWorld!"})
}

func (a *App) handleLogin(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "login.go.html", gin.H{
			"username": form.Username,
			"remember": form.Remember,
			"invalid":  true,
			"message":  "Invalid username or password",
		})
		return
	}
	user, err := a.users.GetUser(form.Username)
	if err != nil {
		log.Printf("Failed to get user %s: %v", form.Username, err)
		c.HTML(http.StatusBadRequest, "login.go.html", gin.H{
			"username": form.Username,
			"remember": form.Remember,
			"invalid":  true,
			"message":  "Invalid username or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		log.Printf("Invalid password")
		c.HTML(http.StatusBadRequest, "login.go.html", gin.H{
			"username": form.Username,
			"remember": form.Remember,
			"invalid":  true,
			"message":  "Invalid username or password",
		})
		return
	}
	cookieAge := 0
	if form.Remember {
		cookieAge = 3600
	}
	c.SetCookie("username", form.Username, cookieAge, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
