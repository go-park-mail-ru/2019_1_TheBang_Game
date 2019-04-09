package main

import (
	"BangGame/config"
	"BangGame/pkg/app"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
