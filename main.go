package main

import (
	"BangGame/pkg/room"

	"BangGame/pkg/game"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	game := game.Game{}
	game.NewGame()

	router.GET("/room", room.RoomsListHandle)

	router.Run(":8081")
}
