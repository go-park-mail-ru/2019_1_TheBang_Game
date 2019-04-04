package game

import (
	"BangGame/config"
	"BangGame/pkg/room"
)

type Game struct {
	MaxRoomsCount uint                `max_rooms_count`
	Rooms         map[uint]*room.Room `rooms`
	RoomsCount    uint                `rooms_count`
}

func (g *Game) NewGame() *Game {
	config.Logger.Infow("NewGame",
		"msg", "Game was created",
	)

	return &Game{
		Rooms:         make(map[uint]*room.Room),
		MaxRoomsCount: config.MaxRoomsInGame,
	}
}

func (g *Game) NewRoom() (*Room, error) {
	if g.RoomsCount == g.MaxRoomsCount {
		config.Logger.Warnw("NewRoom",
		"msg", "Rooms limit")

		return nil, ErrorMaxRoomsLimit
	}

	facker, _ := faker.New("en")
	roomName := facker.Word()

	id := g.RoomsCount + 1

	config.Logger.Infow("NewRoom",
	"msg", fmt.SprintF("New room [id:%v, name:%v] was created", id, roomName))

	return &room.Room{
		Id: Ro
	}, nil
}

func (g *Game) DeleteRoom(id uint) {

}
