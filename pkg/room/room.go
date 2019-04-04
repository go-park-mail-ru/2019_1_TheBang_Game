package room

type Room struct {
	Id           uint      `json:"id"`
	Name         string    `json:"room"`
	MaxPlayers   string    `json: max_players`
	PlayersCount uint      `players_count`
	Players      []*Player `json:"players"`
}
