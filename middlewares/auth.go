package middlewares

import (
	"log"
	"net/http"
)

func JwtAuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware passed")
		next(w, r)
	}
}