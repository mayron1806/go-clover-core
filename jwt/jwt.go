package jwt

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

type Subject map[string]any
type JWTClaims struct {
	jwt.RegisteredClaims
	Subject Subject `json:"sub"`
}
type JWTTokenOptions struct {
	Duration time.Duration `env:"JWT_DURATION" default:"15m"`
	Issuer   string        `env:"JWT_ISSUER"`
	Secret   string        `env:"JWT_SECRET"`
	Method   string        `env:"JWT_METHOD" default:"HS256"`
}

// DefaultTokenOptions returns default token options
func DefaultTokenOptions() (JWTTokenOptions, error) {
	envLoader := config.NewEnvLoader[JWTTokenOptions]()
	env, err := envLoader.LoadEnv()
	if err != nil {
		return JWTTokenOptions{}, err
	}
	return *env, nil
}
func GenerateToken(subject Subject, opts JWTTokenOptions) (string, error) {
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(opts.Duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Subject: subject,
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
func ParseToken(token string, opts JWTTokenOptions) (*JWTClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(opts.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
