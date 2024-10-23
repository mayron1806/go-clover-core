package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/mayron1806/go-clover-core"
)

const (
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)

func GzipMiddleware(level int) clover.IMiddleware {
	return func(next clover.HandlerFunc) clover.HandlerFunc {
		return func(ctx *clover.Context) {
			if strings.Contains(ctx.Request.Header.Get("Accept-Encoding"), "gzip") {
				ctx.Writer.Header().Set("Content-Encoding", "gzip")
				gz, err := gzip.NewWriterLevel(ctx.Writer, level)
				if err != nil {
					http.Error(ctx.Writer, "Failed to apply compression", http.StatusInternalServerError)
					return
				}
				defer gz.Close()

				gzrw := gzipResponseWriter{Writer: gz, ResponseWriter: ctx.Writer}
				next(&clover.Context{
					Request: ctx.Request,
					Writer:  gzrw,
					Params:  ctx.Params,
				})
			} else {
				next(ctx)
			}
		}
	}
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (gzw gzipResponseWriter) Write(b []byte) (int, error) {
	return gzw.Writer.Write(b)
}
