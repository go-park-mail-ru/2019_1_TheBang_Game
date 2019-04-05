package middleware

import (
	"net/http"
)

type urlMehtod struct {
	URL    string
	Method string
}

var ignorCheckAuth = map[urlMehtod]bool{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check := urlMehtod{URL: r.URL.Path, Method: r.Method}
		if ok := ignorCheckAuth[check]; !ok {
			if r.Method == "OPTIONS" {
				return
			}

			if _, ok := CheckTocken(r); !ok {
				w.WriteHeader(http.StatusUnauthorized)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", config.FrontentDst)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}
