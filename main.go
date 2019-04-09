package main

import (
	"BangGame/config"
	"BangGame/pkg/app"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CorsMiddleware)

	router.GET("/room", app.RoomsListHandle)
	router.POST("/room", app.CreateRoomHandle)
	router.GET("/room/:id", app.ConnectRoomHandle)

	router.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			// "state": game.NewMap(),
		})
	})

	router.Run(":" + config.PORT)
}

func CorsMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", config.FrontentDst)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	c.Next()
}
