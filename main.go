package main

func main() {
	//gin.SetMode(gin.ReleaseMode)
	app := NewApp()
	app.Run()
}
