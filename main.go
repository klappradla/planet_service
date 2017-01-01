package main

import (
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/klappradla/planet_service/handler"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/planets", handler.Planets)

	router.Run("localhost:3000")
}
