package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mayron1806/go-clover-core/config"
)

var SigningMethod map[string]*jwt.SigningMethodHMAC = map[string]*jwt.SigningMethodHMAC{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
}

type TokenOptions struct {
	Duration time.Duration `env:"JWT_DURATION" default:"15m"`
	Issuer   string        `env:"JWT_ISSUER"`
	Secret   string        `env:"JWT_SECRET"`
	Method   string        `env:"JWT_METHOD" default:"HS256"`
}
type JWTService struct {
}

// DefaultTokenOptions returns default token options
func DefaultTokenOptions() (TokenOptions, error) {
	envLoader := config.NewEnvLoader[TokenOptions]()
	env, err := envLoader.LoadEnv()
	if err != nil {
		return TokenOptions{}, err
	}
	return *env, nil
}
func (s *JWTService) GenerateToken(payload interface{}, opts TokenOptions) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(opts.Duration)),
		Subject:   payload.(string),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	if opts.Issuer != "" {
		claims.Issuer = opts.Issuer
	}
	token := jwt.NewWithClaims(
		SigningMethod[opts.Method],
		claims,
	)
	res, err := token.SignedString([]byte(opts.Secret))
	return res, err
}
func (s *JWTService) ParseToken(token string, opts TokenOptions) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(opts.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.RegisteredClaims); ok && tokenClaims.Valid {
			return claims.Subject, nil
		}
	}
	return "", err
}

func NewJWTService() *JWTService {
	return &JWTService{}
}
