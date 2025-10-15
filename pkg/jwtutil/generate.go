package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
	"resedist/pkg/config"
	"time"
)

type Claims struct {
	ID         uint
	Type       string
	ClientType string
	jwt.RegisteredClaims
}

// GenerateRefreshToken creates a refresh token for a given user ID and client type
func GenerateRefreshToken(userID uint, clientType string) (string, error) {
	claims := &Claims{
		ID:         userID,
		Type:       "refresh",
		ClientType: clientType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Get().Jwt.RefreshDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get().Jwt.Secret))
}

// GenerateAccessToken creates an access token for a given user ID and client type
func GenerateAccessToken(userID uint, clientType string) (string, error) {
	claims := &Claims{
		ID:         userID,
		Type:       "access",
		ClientType: clientType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Get().Jwt.AccessDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get().Jwt.Secret))
}
