package game

import (
	"BangGame/config"
	"BangGame/pkg/room"
	"fmt"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/manveru/faker"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  config.SocketReadBufferSize,
	WriteBufferSize: config.SocketWriteBufferSize,
}

var GameInst = NewGame()

type Game struct {
	MaxRoomsCount uint                `json:"max_rooms_count"`
	Rooms         map[uint]*room.Room `json:"rooms"`
	RoomsCount    uint                `json:"rooms_count"`
	locker        sync.Mutex
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

func checkRoomID(id string) bool {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	GameInst.locker.Lock()
	defer GameInst.locker.Unlock()

	if _, ok := GameInst.Rooms[uint(ID)]; !ok {
		return false
	}

	return true
}

func (g *Game) WrappedRoomsList() []room.RoomWrap {
	g.locker.Lock()
	defer g.locker.Unlock()

	wraps := []room.RoomWrap{}

	for id := range g.Rooms {
		roomNode, _ := g.Rooms[id]
		wrap := room.WrapedRoom(roomNode)

		wraps = append(wraps, wrap)
	}

	return wraps
}

func (g *Game) RoomsList() []*room.Room {
	g.locker.Lock()
	defer g.locker.Unlock()

	rooms := []*room.Room{}
	for _, room := range g.Rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

func (g *Game) Room(id uint) (*room.Room, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	room, ok := g.Rooms[id]
	if !ok {
		return nil, ErrorRoomNotFound
	}

	return room, nil
}

func (g *Game) WrappedRoom(id uint) (room.RoomWrap, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	gameRoom, ok := GameInst.Rooms[id]
	if !ok {
		return room.RoomWrap{}, ErrorRoomNotFound
	}

	wrap := room.WrapedRoom(gameRoom)

	return wrap, nil
}

// Изменить способ получения id комнаты, возможны коллизии
func (g *Game) NewRoom() (room.RoomWrap, error) {
	g.locker.Lock()
	defer g.locker.Unlock()

	if g.RoomsCount == g.MaxRoomsCount {
		config.Logger.Warnw("NewRoom",
			"msg", "Rooms limit")

		return room.RoomWrap{}, ErrorMaxRoomsLimit
	}

	facker, _ := faker.New("en")
	roomName := facker.Name()

	id := g.RoomsCount + 1
	g.Rooms[id] = &room.Room{
		Id:         id,
		Name:       roomName,
		MaxPlayers: config.MaxPlayersInRoom,
		Register:   make(chan *room.Player),
		Unregister: make(chan *room.Player),
		Players:    make(map[*room.Player]interface{}),
	}
	g.RoomsCount++

	// Запуск комнаты
	go g.Rooms[id].RunRoom()

	config.Logger.Infow("NewRoom",
		"msg", fmt.Sprintf("New room [id:%v, name:%v] was created", id, roomName))

	wrap := room.WrapedRoom(g.Rooms[id])
	return wrap, nil
}

func (g *Game) DeleteRoom(id uint) {
	g.locker.Lock()
	defer g.locker.Unlock()

}
