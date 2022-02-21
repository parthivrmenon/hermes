package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/parthivrmenon/hermes/app"
	"github.com/parthivrmenon/hermes/app/link"
)

func main() {
	gin.SetMode(gin.DebugMode)

	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	repo := link.NewRedisRepository(client)
	linkHandler := link.NewHandler(repo)

	router := gin.Default()
	app.InitializeRoutes(router, linkHandler)
	router.Run(":80")
}
