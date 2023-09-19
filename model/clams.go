package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaim struct {
	UserID string
	Role   string
	jwt.StandardClaims
}