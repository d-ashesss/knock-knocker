package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"path/filepath"
)

const pwd = "$2a$10$S82nUqYWtQ52lO5AJCKFiO..X0los7LU9oOx.CbkhoHVkf4vo7EC6" // 123

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Remember bool   `form:"remember"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.HTMLRender = loadTemplates("templates")
	s.StaticFS("/static", http.Dir("static"))

	s.GET("/", func(c *gin.Context) {
		_, err := c.Cookie("username")
		if err != nil {
			c.HTML(http.StatusOK, "login.gohtml", gin.H{})
			return
		}
		c.HTML(http.StatusOK, "index.gohtml", gin.H{"message": "HellWorld!"})
	})
	s.POST("/", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "login.gohtml", gin.H{
				"username": form.Username,
				"remember": form.Remember,
				"invalid":  true,
				"message":  "Invalid username or password",
			})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(form.Password)); err != nil {
			c.HTML(http.StatusBadRequest, "login.gohtml", gin.H{
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
	})
	log.Printf("Starting server at http://localhost:8181")
	if err := s.Run(":8181"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
