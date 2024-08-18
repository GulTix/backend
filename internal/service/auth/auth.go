package auth

import (
	"backend/internal/entity"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *serviceImpl) GenerateToken(user entity.User) (*string, error) {
	expiredString := os.Getenv("JWT_EXPIRED_TIME")
	jwtSecret := []byte (os.Getenv("JWT_SECRET"))
	signMethod := jwt.SigningMethodHS256

	expiredTime, err := time.ParseDuration(expiredString)

	if err != nil {
		log.Println("[Auth][GenerateToken][ParseDuration] Error Parsing Expired ENV")
		return nil, err
	}

	claims := TokenClaims{
		jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			Subject:   user.Id,
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user,
	}

	token := jwt.NewWithClaims(signMethod, claims)

	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Printf("[Auth][GenerateToken][SignedString] Error Signing Token +%v\n", err)
		return nil, err
	}

	return &signedToken, nil
}
