package room

import (
	"BangGame/api"
	"BangGame/config"
	"fmt"
	"sync"
	"time"
)

type RoomWrap struct {
	Id           uint     `json:"id"`
	Name         string   `json:"room"`
	MaxPlayers   uint     `json:"max_players"`
	PlayersCount uint     `json:"players_count"`
	Players      []Player `json:"players"`
}

func WrapedRoom(room *Room) RoomWrap {
	room.locker.Lock()
	defer room.locker.Unlock()

	palyers := []Player{}
	for player := range room.Players {
		palyers = append(palyers, *player)
	}

	wrap := RoomWrap{
		Id:           room.Id,
		Name:         room.Name,
		MaxPlayers:   room.MaxPlayers,
		PlayersCount: room.PlayersCount,
		Players:      palyers,
	}

	return wrap
}

// Подумать о том, как это все таки будет передаваться json-ом
type Room struct {
	Id           uint                    `json:"id"`
	Name         string                  `json:"room"`
	MaxPlayers   uint                    `json:"max_players"`
	PlayersCount uint                    `json:"players_count"`
	Players      map[*Player]interface{} `json:"players"`
	Register     chan *Player            `json:"-"`
	Unregister   chan *Player            `json:"-"`
	Broadcast    chan api.SocketMsg      `json:"-"`
	Closer       chan struct{}           `json:"-"`
	locker       sync.Mutex              `json:"-"`
}

func (r *Room) Conection(player *Player) {
	if r.PlayersCount == r.MaxPlayers {
		player.Conn.WriteJSON(api.TooManyPlayersMsg)
		player.Conn.Close()
	}

	r.locker.Lock()
	r.Players[player] = nil
	r.locker.Unlock()

	r.PlayersCount++

	config.Logger.Infow("Conection",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was connected to room [id: %v, name: %v]",
			player.Id, player.Nickname, r.Id, r.Name))
}

func (r *Room) Disconection(player *Player) {
	r.locker.Lock()
	delete(r.Players, player)
	r.locker.Unlock()

	player.Conn.Close()

	r.PlayersCount--

	config.Logger.Infow("Conection",
		"msg", fmt.Sprintf("Player [id: %v, nick: %v] was disconnected from room [id: %v, name: %v]",
			player.Id, player.Nickname, r.Id, r.Name))
}

// Заглушка для рассылки
func (r *Room) SnapShot() {
	for player := range r.Players {
		player.In <- api.SocketMsg{
			Type: "test",
			Data: struct {
				Msg string
			}{Msg: "test"},
		}
	}
}

func (r *Room) Distribution(msg api.SocketMsg) {
	for player := range r.Players {
		player.In <- msg
	}
}

func (r *Room) RunRoom() {
	config.Logger.Infow("RunRoom",
		"msg", fmt.Sprintf("Room  [id: %v name: %v] opened", r.Id, r.Name))

	defer config.Logger.Infow("RunRoom",
		"msg", fmt.Sprintf("Room [id: %v name: %v] closed", r.Id, r.Name))

	ticker := time.NewTicker(config.RoomTickTime)
	defer ticker.Stop()

Loop:
	for {
		select {
		case player := <-r.Register:
			r.Conection(player)

		case player := <-r.Unregister:
			r.Disconection(player)

		case msg := <-r.Broadcast:
			r.Distribution(msg)

		case <-ticker.C:
			r.SnapShot()

		case <-r.Closer:
			break Loop
		}
	}

	// Тут наверно должен бвть дисконект всех, кто все таки остался в комнате
}
