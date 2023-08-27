package services

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
}

func (service *JwtService) Sign(data jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func (service *JwtService) ValidateToken(token string, claims jwt.MapClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
}

func NewJwtService() *JwtService {
	return &JwtService{}
}
