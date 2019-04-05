package room

import "sync"

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
	register     chan *Player
	unregister   chan *Player
	broadcast    chan []byte
	locker       sync.Mutex
}

func (r *Room) RunRoom() {
	for {
		select {
		case player := <-r.register:
			// ToDo подумать, может все таки bool, а не пустой интерфейс
			r.Players[player] = nil

		case player := <-r.unregister:
			if _, ok := r.Players[player]; ok {
				delete(r.Players, player)
				// Todo закрывать все необходимое в этом цикле
				// close(player.send)
			}
		}

		// ToDo место под снапшот
	}
}
