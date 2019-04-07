package api

const RoomState = "room"

var TooManyPlayersMsg = SocketMsg{
	Type: "error",
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "Too many players in this room",
	},
}
