package game

import (
	"BangGame/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "can not upgrade to websocket",
		})
	}

	room, _ := GameInst.WrappedRoom(id)
	msg := api.SockMsg{
		Status: "room",
		Data:   room,
	}

	websocket.WriteJSON(conn, msg)
}
