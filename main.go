package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

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
			c.HTML(http.StatusBadRequest, "login.gohtml", gin.H{"message": "Invalid username or password"})
			return
		}
		c.SetCookie("username", form.Username, 3600, "/", "localhost", false, true)
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
