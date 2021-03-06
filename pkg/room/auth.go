package room

import (
	"fmt"
	"log"
	"net/http"

	"BangGame/config"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Id       float64 `json:"id"`
	Nickname string  `json:"nickname"`
	PhotoURL string  `json:"photo_url"`

	jwt.StandardClaims
}

func TokenFromCookie(r *http.Request) *jwt.Token {
	cookie, _ := r.Cookie(config.CookieName)
	tokenStr := cookie.Value
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})
	return token
}

func InfoFromCookie(token *jwt.Token) (userInfo UserInfo, status int) {
	userInfo = UserInfo{}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userInfo.Id = claims["id"].(float64)
		userInfo.Nickname = claims["nickname"].(string)
		userInfo.PhotoURL = claims["photo_url"].(string)
	} else {
		status = http.StatusInternalServerError
		config.Logger.Warnw("NicknameFromCookie",
			"warn", "Error with parsing token's claims")

		return userInfo, status
	}

	return userInfo, http.StatusOK
}

func CheckTocken(r *http.Request) (token *jwt.Token, ok bool) {
	// дебажу
	fmt.Println("!!!!")
	fmt.Println("!!!!", r.Method, "!!!!")
	fmt.Println("!!!!", r.Cookies(), "!!!!")
	fmt.Println("!!!!")
	// дебажу

	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		config.Logger.Warnw("CheckTocken",
			"warn", err.Error())
		return nil, false
	}

	tokenStr := cookie.Value

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.SECRET, nil
	})
	if err != nil {
		log.Printf("Error with check tocken: %v", err.Error())

		return nil, false
	}

	if !token.Valid {
		log.Printf("%v use faked cookie: %v\n", r.RemoteAddr, err.Error())

		return nil, false
	}

	return token, true
}
