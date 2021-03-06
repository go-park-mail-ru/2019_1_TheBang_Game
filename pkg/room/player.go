package room

import (
	"BangGame/api"
	"BangGame/config"
	"BangGame/pkg/auth"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type UserInfo struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`
}

type Player struct {
	Id       float64            `json:"id"`
	Nickname string             `json:"nickname"`
	PhotoURL string             `json:"photo_url"`
	Conn     *websocket.Conn    `json:"-"`
	In       chan api.SocketMsg `json:"-"`
	Out      chan api.SocketMsg `json:"-"`
	Room     *Room              `json:"-"`
}

func (p *Player) Reading() {
	ticker := time.NewTicker(config.PlayerReadingTickTime)
	defer func() {
		ticker.Stop()
		p.Conn.Close()
	}()

Loop:
	for {
		msg, ok := <-p.In
		if !ok {
			break Loop
		}

		err := p.Conn.WriteJSON(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (p *Player) Writing() {
	go func() {
		for {
			msg := &api.SocketMsg{}
			err := p.Conn.ReadJSON(msg)
			if websocket.IsUnexpectedCloseError(err) {
				p.Room.Unregister <- p
				return
			}

			p.Out <- *msg
		}
	}()

Loop:
	for {
		msg, ok := <-p.Out
		if !ok {
			break Loop
		}

		p.Room.Broadcast <- msg
	}
}

func PlayerFromCtx(ctx *gin.Context, conn *websocket.Conn) *Player {
	info := playerInfoFromCookie(ctx)
	player := &Player{
		Id:       info.Id,
		Nickname: info.Nickname,
		PhotoURL: info.PhotoURL,
		Conn:     conn,
		In:       make(chan api.SocketMsg, config.InOutBuffer),
		Out:      make(chan api.SocketMsg, config.InOutBuffer),
	}

	config.Logger.Infow("PlayerFromCtx",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was initialized", player.Id, player.Nickname))

	return player
}

func playerInfoFromCookie(ctx *gin.Context) UserInfo {
	info, _ := auth.CheckTocken(ctx.Request)

	return UserInfo{
		Id:       info.Id,
		Nickname: info.Nickname,
		PhotoURL: info.PhotoUrl,
	}
}
