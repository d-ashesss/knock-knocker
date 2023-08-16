package main

func main() {
	//gin.SetMode(gin.ReleaseMode)
	appCfg := NewConfig()
	app := NewApp(appCfg)
	app.Run()
}
