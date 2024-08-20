package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				exception.WriteUnauthorizedError(w, "Unauthorized")
				return
			}
			exception.WriteUnauthorizedError(w, "Unauthorized")
			return
		}
		
		// mengambil token value
		tokenString := c.Value
		
		claims := &util.JWTClaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return util.JWT_KEY, nil
		})
		
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				// token invalid
				exception.WriteUnauthorizedError(w, "Unauthorized")
				return
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				// token expired
				exception.WriteUnauthorizedError(w, "Unauthorized, Token expired!")
				return
				} else {
					// Handle other errors
					exception.WriteUnauthorizedError(w, "Unauthorized")
					return
				}
			}
			if !token.Valid {
				exception.WriteUnauthorizedError(w, "Unauthorized")
				return
			}
			
			ctx := context.WithValue(r.Context(), util.TokenKey , tokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
