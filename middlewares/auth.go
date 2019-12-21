package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go-authen/repo"
	"go-authen/utils"
	"net/http"
	"os"
	"strings"
)

// Adapted from https://github.com/hellojebus/go-mux-jwt-boilerplate/blob/master/middleware/verifyJWT.go
func JwtAuthenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		// Check if auth header is present
		if len(tokenStr) == 0 {
			msg := utils.NewMessage(false, 5, "no authorization header", nil)
			utils.Respond(w, http.StatusUnauthorized, msg)
			return
		}

		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

		// Parse token
		claims, token, claimErr := parseToken(tokenStr)
		if claimErr != nil || !claims.Valid {
			msg := utils.NewMessage(false, 5, "invalid token", nil)
			utils.Respond(w, http.StatusUnauthorized, msg)
			return
		}

		// Verified token, pass to the next handler
		ctx := context.WithValue(r.Context(), "user", token.UserId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

func parseToken(tokenStr string) (*jwt.Token, *repo.Token, error) {
	parsedToken := &repo.Token{}
	claims, err := jwt.ParseWithClaims(tokenStr, parsedToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("token_secret")), nil
	})
	return claims, parsedToken, err
}
