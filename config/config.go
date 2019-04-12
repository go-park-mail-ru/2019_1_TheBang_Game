package config

import (
	"os"
	"time"
)

const (
	MaxPlayersInRoom uint = 2 // поправить и вернуть обратно
	MaxRoomsInGame   uint = 10
	RoomTickTime          = 5 * time.Second
	// GameTickTime               = 1 * time.Second // fps стоит обсуждений
	PlayerWritingTickTime = 1 * time.Second
	PlayerReadingTickTime = 1 * time.Second

	WriteDeadline = 10 * time.Second
	// ReadingWait = 10 * time.Second
)

var Logger = NewGlobalLogger()

var (
	SocketReadBufferSize        = 1024
	SocketWriteBufferSize       = 1024
	MaxMessageSize        int64 = 512
	InOutBuffer                 = 10
)

var (
	CookieName  = "bang_token"
	SECRET      = getSecret()
	PORT        = getPort()
	FrontentDst = getFrontDest()
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

func getFrontDest() string {
	dst := os.Getenv("FrontentDst")
	if dst == "" {
		Logger.Warn("There is no FrontentDst!")
		dst = "http://127.0.0.1:3000"
	}
	return dst
}
