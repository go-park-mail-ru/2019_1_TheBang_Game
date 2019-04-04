package main

import (
	"BangGame/pkg/room"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/room", room.RoomsListHandle)

	router.Run(":8081")
}
