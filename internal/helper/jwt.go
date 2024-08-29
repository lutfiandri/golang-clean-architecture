package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lutfiandri/golang-clean-architecture/internal/entity"
	"github.com/lutfiandri/golang-clean-architecture/internal/model"
)

func GenerateJwt(user entity.User, secret string) (string, error) {
	claims := model.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		User: model.JwtUser{
			ID:   user.ID,
			Role: user.RoleID,
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokens.SignedString([]byte(secret))
}

func ParseJwt(tokenString, secret string) (model.JwtClaims, error) {
	var claims model.JwtClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}

	return claims, err
}
