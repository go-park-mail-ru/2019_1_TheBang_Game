package main

import (
	"BangGame/config"
	"BangGame/pkg/app"
	"BangGame/pkg/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware,
		middleware.AuthMiddleware)

	router.GET("/room", app.RoomsListHandle)
	router.POST("/room", app.CreateRoomHandle)
	router.GET("/room/:id", app.ConnectRoomHandle)

	return router
}

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("FrontenDest: %v", config.FrontentDst))
	config.Logger.Info(fmt.Sprintf("PORT: %v", config.PORT))

	router := setUpRouter()

	router.Run(":" + config.PORT)
}
