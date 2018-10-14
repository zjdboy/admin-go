package utils

import (
	"admin-go/defs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(username string) (string, error) {
	expire := time.Now().Add(3 * time.Hour)

	claims := defs.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := claim.SignedString([]byte(defs.SECRETKEY))

	return token, err
}

func ParseToken(token string) (*defs.Claims, error) {
	claim, err := jwt.ParseWithClaims(token, &defs.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(defs.SECRETKEY), nil
	})

	if claim != nil {
		if claims, ok := claim.Claims.(*defs.Claims); ok && claim.Valid {
			return claims, nil
		}
	}

	return nil, err
}
