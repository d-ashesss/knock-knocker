package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Remember bool   `form:"remember"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.LoadHTMLGlob("templates/*")
	s.StaticFS("/static", http.Dir("static"))
	s.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{"message": "HellWorld!"})
	})
	s.POST("/", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			log.Print(err)
			c.HTML(http.StatusBadRequest, "index.gohtml", gin.H{"message": "Invalid username or password"})
			return
		}
		log.Print(form)
		log.Printf("login attempt!")
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	if err := s.Run(":8181"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
