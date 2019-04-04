package room

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	PhotoURL string `json:"photo_url"`
	conn     *websocket.Conn
}
