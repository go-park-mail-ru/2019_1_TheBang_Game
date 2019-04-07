package room

import (
	"BangGame/api"
	"BangGame/config"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Player struct {
	Id       uint               `json:"id"`
	Nickname string             `json:"nickname"`
	PhotoURL string             `json:"photo_url"`
	Conn     *websocket.Conn    `json:"-"`
	In       chan api.SocketMsg `json:"-"`
	Out      chan api.SocketMsg `json:"-"`
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
			log.Println(err.Error())
			break Loop
		}
	}

	// Loop:
	// 	for {
	// 		select {
	// 		case msg, ok := <-p.In:
	// 			p.Conn.SetWriteDeadline(time.Now().Add(config.WriteDeadline))
	// 			if !ok {
	// 				// ToDo подумать о сообщении для дисконекта
	// 				p.Conn.WriteMessage(websocket.CloseMessage, []byte{})
	// 				return
	// 			}

	// 			err := p.Conn.WriteJSON(msg)
	// 			if err != nil {
	// 				log.Println(err.Error())
	// 				break Loop
	// 			}

	// 		case <-ticker.C:
	// 			log.Println(p.Nickname, "Reading")
	// 			p.In <- api.SocketMsg{
	// 				Type: "test",
	// 				Data: struct {
	// 					Msg string
	// 				}{Msg: "test"},
	// 			}
	// 		}
	// 	}
}

func (p *Player) Writing() {
	// // p.Conn.SetReadLimit(config.MaxMessageSize)
	// // p.Conn.SetPongHandler(func(string) error { p.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// ticker := time.NewTicker(config.PlayerReadingTickTime)
	// defer func() {
	// 	ticker.Stop()
	// 	p.Conn.Close()
	// }()

	// for {
	// 	select {
	// 	case msg, ok := <-p.Out:

	// 	}
	// }
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
		In:       make(chan api.SocketMsg, config.InOutBuffer),
		Out:      make(chan api.SocketMsg, config.InOutBuffer),
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
