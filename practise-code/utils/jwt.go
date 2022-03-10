package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"practise-code/global"
	jwtModel "practise-code/model/http/jwt"
	"practise-code/model/sql/user"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func GetToken(user user.User) (string, error) {
	claim := jwtModel.BaseClaims{
		ID: user.DefaultField.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// 这里需要明确指出时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(global.CONFIG.JWT.ExpiresTime))), // 过期时间
			IssuedAt: jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),     // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	return token.SignedString([]byte(global.CONFIG.JWT.SigningKey))
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	}
}

func ParseToken(tokens string) (*jwtModel.BaseClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &jwtModel.BaseClaims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
		return nil, errors.New("couldn't handle this token")
	}
	if claims, ok := token.Claims.(*jwtModel.BaseClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
