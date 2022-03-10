package jwt

import "github.com/golang-jwt/jwt/v4"

type BaseClaims struct {
	ID          uint
	Username    string
	jwt.RegisteredClaims
}
