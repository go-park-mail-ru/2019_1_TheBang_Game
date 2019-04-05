package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkRoomID(id string) bool {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	if _, ok := GameInst.Rooms[uint(ID)]; !ok {
		return false
	}

	return true
}

func ConnectRoomHandle(c *gin.Context) {
	id := c.Param("id")
	if ok := checkRoomID(id); !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ok := c.IsWebsocket(); !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
