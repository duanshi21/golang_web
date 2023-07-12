package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var stSigningKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCostClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string) (string, error) {
	iJwtCostClaims := JwtCostClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCostClaims)
	return token.SignedString(stSigningKey)
}

func ParseToken(tokenStr string) (JwtCostClaims, error) {
	iJitCostClaims := JwtCostClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, &iJitCostClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return JwtCostClaims{}, err
}

// IsTokenValid 判断token是否有效
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
