package room

type Room struct {
	Id           uint   `json:"id"`
	Name         string `json:"room"`
	MaxPlayers   string
	Players      []*Player `json:"players"`
	PlayersCount uint      `players_count`
}
