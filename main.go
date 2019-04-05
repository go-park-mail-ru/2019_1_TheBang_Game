package main

import (
	"BangGame/pkg/game"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/room", game.RoomsListHandle)
	router.POST("/room", game.CreateRoomHandle)
	router.GET("/room/:id", game.ConnectRoomHandle)

	router.Run(":8081")

}
