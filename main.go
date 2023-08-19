package main

import (
	"github.com/d-ashesss/knock-knocker/datastore"
	"github.com/d-ashesss/knock-knocker/users"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	dbCfg, err := datastore.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}
	db, err := datastore.Open(dbCfg)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	usersSrv, err := users.NewService(db)
	if err != nil {
		log.Fatalf("Failed to create users service: %v", err)
	}

	appCfg := NewConfig()
	app := NewApp(appCfg, usersSrv)
	app.Run()
}
