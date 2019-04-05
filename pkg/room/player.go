package room

import (
	"BangGame/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Player struct {
	Id       uint            `json:"id"`
	Nickname string          `json:"nickname"`
	PhotoURL string          `json:"photo_url"`
	Conn     *websocket.Conn `json:"-"`
	In       chan []byte     `json:"-"`
	Out      chan []byte     `json:"-"`
}

type UserInfo struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	PhotoURL string `json:"photo_url"`
}

func PlayerFromCtx(ctx *gin.Context, conn *websocket.Conn) *Player {
	info := playerInfoFromCookie(ctx)
	player := &Player{
		Id:       info.Id,
		Nickname: info.Nickname,
		PhotoURL: info.PhotoURL,
		Conn:     conn,
		In:       make(chan []byte, config.InOutBuffer),
		Out:      make(chan []byte, config.InOutBuffer),
	}

	config.Logger.Infow("PlayerFromCtx",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was initialized", player.Id, player.Nickname))

	return player
}

// ToDo просто заглушка на первое время
func playerInfoFromCookie(ctx *gin.Context) UserInfo {
	return UserInfo{
		Id:       1,
		Nickname: "test",
		PhotoURL: "test",
	}
}
