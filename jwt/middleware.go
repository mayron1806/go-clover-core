package jwt

import (
	"net/http"
	"strings"

	"github.com/mayron1806/go-clover-core"
)

type TokenLocation string

const (
	Header TokenLocation = "header"
	Query  TokenLocation = "query"
	Cookie TokenLocation = "cookie"
)

type JWTMiddlewareOptions struct {
	JWTOpts       JWTTokenOptions
	TokenLocation TokenLocation
}

func DefaultJWTMiddlewareOptions() (*JWTMiddlewareOptions, error) {
	defaultJWTOptions, err := DefaultTokenOptions()
	if err != nil {
		return nil, err
	}
	return &JWTMiddlewareOptions{
		JWTOpts:       defaultJWTOptions,
		TokenLocation: Header,
	}, nil
}
func JWTMiddleware(opts *JWTMiddlewareOptions) clover.IMiddleware {
	return func(next clover.HandlerFunc) clover.HandlerFunc {
		return func(ctx *clover.Context) {
			var token string
			switch opts.TokenLocation {
			case Header:
				token = ctx.Request.Header.Get("Authorization")
			case Query:
				token = ctx.Request.URL.Query().Get("token")
			case Cookie:
				cookie, err := ctx.Request.Cookie("token")
				if err != nil {
					ctx.ResponseJSON(http.StatusUnauthorized, clover.H{"error": "Token is missing"})
					return
				}
				token = cookie.Value
			default:
				ctx.ResponseJSON(http.StatusUnauthorized, clover.H{"error": "Token is missing"})
				return
			}

			if token == "" {
				ctx.ResponseJSON(http.StatusUnauthorized, clover.H{"error": "Token is missing"})
				return
			}
			// Remove the "Bearer " prefix if it exists
			tokenString := strings.TrimPrefix(token, "Bearer ")
			if tokenString == token {
				ctx.ResponseJSON(http.StatusUnauthorized, clover.H{"error": "Authorization header format must be Bearer {token}"})
				return
			}

			claims, err := ParseToken(tokenString, opts.JWTOpts)
			if err != nil {
				ctx.ResponseJSON(http.StatusUnauthorized, clover.H{"error": "Invalid token: " + err.Error()})
				return
			}
			ctx.Set("claims", claims)
			next(ctx)
		}
	}
}
