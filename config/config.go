package config

import "os"

const (
	MaxPlayersInRoom uint = 4
	MaxRoomsInGame   uint = 10
)

var Logger = NewGlobalLogger()

var (
	SocketReadBufferSize  = 1024
	SocketWriteBufferSize = 1024
)

var InOutBuffer = 10

var (
	CookieName = "bang_token"
	SECRET     = getSecret()
	PORT       = getPort()
)

func getSecret() []byte {
	secret := []byte(os.Getenv("SECRET"))
	if string(secret) == "" {
		Logger.Warn("There is no SECRET!")
		secret = []byte("secret")
	}

	return secret
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		Logger.Warn("There is no PORT!")
		port = "8081"
	}
	return port
}
