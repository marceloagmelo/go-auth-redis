package main

import (
	"github.com/marceloagmelo/go-auth-redis/logger"

	"github.com/marceloagmelo/go-auth-redis/app"
	"github.com/marceloagmelo/go-auth-redis/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	logger.Info.Println("Listen 8080...")
	app.Run(":8080")
}
