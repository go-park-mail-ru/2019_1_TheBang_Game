package test

import (
	"BangGame/pkg/app"
	"BangGame/pkg/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestCreateRoomHandle(t *testing.T) {
	router := setUpRouter()
	path := "/room"
	method := "POST"

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("TestCreateRoomHandle, room was not created: expected %v, have %v!\n",
			http.StatusCreated, rr.Code)
	}
}
