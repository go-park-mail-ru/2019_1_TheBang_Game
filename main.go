package main

import (
	"BangGame/pkg/game"

	"github.com/gin-gonic/gin"
)

type point struct {
	A uint
}

func main() {
	router := gin.Default()

	router.GET("/room", game.RoomsListHandle)

	router.Run(":8081")

}
