package model

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	User JwtUser `json:"user"`
}

type JwtUser struct {
	ID   uint `json:"id"`
	Role uint `json:"role"`
}
