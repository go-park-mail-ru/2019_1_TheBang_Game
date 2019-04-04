package room

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
}

func (r *Room) RunRoom() {
	for {
		select {
		case player := <-r.register:
			// ToDo подумать, может все таки bool, а не пустой интерфейс
			r.Players[player] = nil

		case player := <-r.register:
			if _, ok := r.Players[player]; ok {
				delete(r.Players, player)
				// Todo закрывать все необходимое в этом цикле
				// close(player.send)
			}
		}

		// ToDo место под снапшот
	}
}
