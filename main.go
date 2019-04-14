package main

import (
	"BangGame/config"
	"BangGame/pkg/app"
	"fmt"
	"net/http"
	"BangGame/pkg/room"



	"github.com/gin-gonic/gin"
)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("PORT: %v", config.PORT))

	router := gin.Default()
	router.Use(CorsMiddleware, AuthMiddleware)

	router.GET("/room", app.RoomsListHandle)
	router.POST("/room", app.CreateRoomHandle)
	router.GET("/room/:id", app.ConnectRoomHandle)

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

func AuthMiddleware(c *gin.Context) {
	_, ok := room.CheckTocken(c.Request)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("Bad token")

		return
	}

	c.Next()
}
