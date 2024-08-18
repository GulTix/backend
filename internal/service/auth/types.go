package auth

import (
	"backend/internal/entity"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Service interface {
		GenerateToken(user entity.User) (*string, error)
	}

	serviceImpl struct {
	}

	TokenClaims struct {
		jwt.Claims
		User entity.User
	}
)
