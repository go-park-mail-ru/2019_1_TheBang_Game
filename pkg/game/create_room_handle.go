package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoomHandle(c *gin.Context) {
	room, err := GameInst.NewRoom()
	if err != nil {
		c.JSONP(http.StatusConflict, room)
		return
	}

	c.JSONP(http.StatusCreated, room)
}
