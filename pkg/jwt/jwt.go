package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vkevv/back-rectifier/pkg/models"
)

var minSecretLen = 128

// New generates new JWT service necessary for auth middleware
func New(algo, secret string, ttlMinutes, minSecretLength int) (Service, error) {
	if minSecretLength > 0 {
		minSecretLen = minSecretLength
	}
	if len(secret) < minSecretLen {
		return Service{}, fmt.Errorf("jwt secret length is %v, which is less than required %v", len(secret), minSecretLen)
	}
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		return Service{}, fmt.Errorf("invalid jwt signing method: %s", algo)
	}

	return Service{
		key: []byte(secret),
		alg: signingMethod,
		ttl: time.Duration(ttlMinutes) * time.Minute,
	}, nil
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	ttl time.Duration

	// JWT signing algorithm
	alg jwt.SigningMethod
}

// ParseToken parses token from Authorization header
func (s Service) ParseToken(authHeader string) (*jwt.Token, error) {
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, fmt.Errorf("Auth header string not valid")
	}

	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if s.alg != token.Method {
			return nil, fmt.Errorf("Algorithm doesn't match")
		}
		return s.key, nil
	})

}

// GenerateToken generates new JWT token and populates it with user data
func (s Service) GenerateToken(u models.User) (string, error) {
	return jwt.NewWithClaims(s.alg, jwt.MapClaims{
		"id":    u.Base.ID,
		"email": u.Email,
		"exp":   time.Now().Add(s.ttl).Unix(),
	}).SignedString(s.key)

}
