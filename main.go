package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/parthivrmenon/hermes/app"
	"github.com/parthivrmenon/hermes/app/link"
)

func main() {
	//setting enivironment
	gin.SetMode(gin.DebugMode)

	//Init REdis
	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	repo := link.NewRedisRepository(client)
	linkHandler := link.NewHandler(repo)

	router := gin.Default()
	app.InitializeRoutes(router, linkHandler)
	router.Run()
}
