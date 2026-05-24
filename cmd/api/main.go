package main

import "tradex/internal/startup"

func main() {
	app := startup.NewApp()
	app.Run()
}
