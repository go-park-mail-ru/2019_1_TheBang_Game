package api

const RoomState = "room_snap_shot"

var TooManyPlayersMsg = SocketMsg{
	Type: "disconection",
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "Too many players in this room",
	},
}

var ConectionMsg = SocketMsg{
	Type: "conection",
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "You was conected to the room",
	},
}
