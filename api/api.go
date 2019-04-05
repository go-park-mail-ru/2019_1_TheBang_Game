package api

type SockMsg struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
