package main

import (
	"BangGame/config"
	"BangGame/pkg/game"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/room", game.RoomsListHandle)
	router.POST("/room", game.CreateRoomHandle)
	router.GET("/room/:id", game.ConnectRoomHandle)

	router.Run(":" + config.PORT)
}
