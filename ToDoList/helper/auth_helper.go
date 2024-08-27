package helper

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"
)

func ExtractUserFromToken(ctx context.Context) (string, string, error) {
	tokenString, ok := ctx.Value(util.TokenKey).(string)
	if !ok {
		return "", "", errors.New("token not found in context")
	}
	claims := &util.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return util.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid token")
	}
	return claims.ID, claims.Username, nil
}
