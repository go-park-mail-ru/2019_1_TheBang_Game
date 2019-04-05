package api

type SocketMsg struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
