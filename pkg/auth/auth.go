package auth

import (
	"BangGame/config"
	"BangGame/pkg/room"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PlayerInfoFromCookie(c *gin.Context) (room.UserInfo, error) {
	cookie, err := c.Cookie(config.CookieName)
	if err != nil {
		return room.UserInfo{}, err
	}

	token, _ := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})

	var claims jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		config.Logger.Warnw("PlayerInfoFromCookie",
			"msg", "Can not convert jwt claims")

		return room.UserInfo{}, fmt.Errorf("Can not convert jwt claims")
	}

	info := room.UserInfo{}
	info.Id = claims["id"].(uint)
	info.Nickname = claims["nickname"].(string)
	info.PhotoURL = claims["photo_url"].(string)

	fmt.Println(info)

	return room.UserInfo{}, nil
}
