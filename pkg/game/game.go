package game

import (
	"BangGame/config"
	"BangGame/pkg/room"
	"fmt"

	"github.com/manveru/faker"
)

var GameInst = NewGame()

type Game struct {
	MaxRoomsCount uint                `json:"max_rooms_count"`
	Rooms         map[uint]*room.Room `json:"rooms"`
	RoomsCount    uint                `json:"rooms_count"`
}

func NewGame() *Game {
	config.Logger.Infow("NewGame",
		"msg", "Game was created",
	)

	return &Game{
		Rooms:         make(map[uint]*room.Room),
		MaxRoomsCount: config.MaxRoomsInGame,
	}
}

func (g *Game) RoomsList() []*room.Room {
	rooms := []*room.Room{}
	for _, room := range g.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

func (g *Game) NewRoom() (*room.Room, error) {
	if g.RoomsCount == g.MaxRoomsCount {
		config.Logger.Warnw("NewRoom",
			"msg", "Rooms limit")

		return nil, ErrorMaxRoomsLimit
	}

	facker, _ := faker.New("en")
	roomName := facker.Name()

	id := g.RoomsCount + 1
	g.Rooms[id] = &room.Room{
		Id:         id,
		Name:       roomName,
		MaxPlayers: config.MaxPlayersInRoom,
	}
	g.RoomsCount++

	config.Logger.Infow("NewRoom",
		"msg", fmt.Sprintf("New room [id:%v, name:%v] was created", id, roomName))

	return g.Rooms[id], nil
}

func (g *Game) DeleteRoom(id uint) {

}
