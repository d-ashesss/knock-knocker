package main

import (
	"github.com/d-ashesss/knock-knocker/datastore"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	dbCfg, err := datastore.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}
	_, err = datastore.Open(dbCfg)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	appCfg := NewConfig()
	app := NewApp(appCfg)
	app.Run()
}
