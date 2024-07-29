package middleware

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/chava.gnolasco/polaris/application/domain/commands"
	"github.com/chava.gnolasco/polaris/infraestructure/env"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(env.GetEnv().JWT_KEY)

/*
JWTAuth is a middleware that checks if the request has a valid JWT token.
*/
func JWTAuth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &commands.JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var isPadminPath bool = strings.Contains(r.URL.Path, "/api/v1/padmin")
		isPadminUser := regexp.MustCompile(`\b(ADM|FS|TS)\b`).MatchString(claims.Roles)

		if isPadminPath && !isPadminUser {
			http.Error(w, "No authorized user", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
