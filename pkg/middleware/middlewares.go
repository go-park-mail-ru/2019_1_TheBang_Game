package middleware

import (
	"BangGame/config"
	"BangGame/pkg/room"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlMehtod struct {
	URL    string
	Method string
}

var ignorCheckAuth = map[urlMehtod]bool{
	urlMehtod{URL: "/room", Method: "POST"}: true,
	urlMehtod{URL: "/room", Method: "GET"}: true,
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
	if c.Request.Method == "OPTIONS" {
		return
	}

	check := urlMehtod{URL: c.Request.URL.Path, Method: c.Request.Method}
	if ok := ignorCheckAuth[check]; !ok {
		_, ok := room.CheckTocken(c.Request)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}
	}

	c.Next()
}
