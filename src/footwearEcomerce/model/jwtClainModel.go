package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
