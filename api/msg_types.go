package api

const RoomState = "room_snap_shot"
const GameState = "game_snap_shot"
const GameStarted = "start_game"
const GameFinish = "finish_game"

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

var GameStartedMsg = SocketMsg{
	Type: GameStarted,
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "Game was started",
	},
}

var GameFinishedMsg = SocketMsg{
	Type: GameFinish,
	Data: struct {
		Msg string `json:"msg"`
	}{
		Msg: "Game was finished",
	},
}
