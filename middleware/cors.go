package middleware

import (
	"net/http"

	"github.com/mayron1806/go-clover-core"
)

func CorsMiddleware(allowedOrigins string, allowedMethods string, allowedHeaders string) clover.IMiddleware {
	return func(next clover.HandlerFunc) clover.HandlerFunc {
		return func(ctx *clover.Context) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			// Verifica se o método é OPTIONS, usado nas preflight requests
			if ctx.Request.Method == http.MethodOptions {
				ctx.Writer.WriteHeader(http.StatusOK)
				return
			}
			next(ctx)
		}
	}
}
