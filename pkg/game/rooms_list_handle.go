package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoomsListHandle(c *gin.Context) {
	rooms := GameInst.RoomsList()
	c.JSONP(http.StatusOK, rooms)
}
