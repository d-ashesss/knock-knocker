package main

import (
	"github.com/d-ashesss/knock-knocker/users"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

type App struct {
	Server *http.Server
	users  users.Service
}

func NewApp(cfg *Config, usersSrv users.Service) *App {
	h := gin.Default()
	app := &App{
		Server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: h,
		},
		users: usersSrv,
	}

	h.HTMLRender = loadTemplates("templates")
	h.StaticFS("/static", http.Dir("static"))

	h.GET("/", app.handleIndex)
	h.POST("/", app.handleLogin)

	return app
}

func (a *App) Run() {
	log.Printf("Starting server at http://localhost%s", a.Server.Addr)
	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.go.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.go.html")
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
