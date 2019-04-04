package room

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Player struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	PhotoURL string `json:"photo_url"`
	conn     *websocket.Conn
	in       chan ([]byte)
	out      chan ([]byte)
}

func PlayerFromCtx(ctx *gin.Context, conn *websocket.Conn) *Player {

}
