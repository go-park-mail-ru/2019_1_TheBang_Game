package game

import "errors"

var (
	ErrorMaxRoomsLimit = errors.New("Rooms limit")
	ErrorRoomNotFound  = errors.New("There id no this room")
)
