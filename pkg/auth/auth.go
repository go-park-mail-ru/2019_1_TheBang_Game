package auth

import (
	"BangGame/config"
	pb "BangGame/pkg/protobuf"
	"context"
	"errors"
	"time"

	"fmt"
	"net/http"

	"google.golang.org/grpc"
)

func GetUserInfo(token string) (user *pb.UserInfo, myerr error) {
	conn, err := grpc.Dial(config.AuthServer, grpc.WithInsecure())
	if err != nil {
		myerr = fmt.Errorf("did not connect: %v", err.Error())

		return
	}
	defer conn.Close()

	client := pb.NewCookieCheckerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.CheckCookie(ctx, &pb.CookieRequest{JwtToken: token})
	if err != nil {
		myerr = fmt.Errorf("could not check: %v", err.Error())

		return
	}

	if !r.Valid {
		myerr = errors.New("invalid cookie")

		return
	}

	user = r.User
	return
}

func CheckTocken(r *http.Request) (info *pb.UserInfo, ok bool) {
	// дебажу
	fmt.Println("!!!!")
	fmt.Println("!!!!", r.Method, "!!!!")
	fmt.Println("!!!!", r.Cookies(), "!!!!")
	fmt.Println("!!!!")
	// дебажу

	cookie, err := r.Cookie(config.CookieName)
	if err != nil {
		config.Logger.Warnw("CheckTocken -> get cookie:",
			"warn", err.Error())

		return
	}

	tokenStr := cookie.Value
	info, err = GetUserInfo(tokenStr)
	if err != nil {
		config.Logger.Warnw("CheckTocken -> GetUserInfo:",
			"warn", err.Error())
	}

	return info, true
}
