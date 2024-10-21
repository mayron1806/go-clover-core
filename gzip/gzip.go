package gzip

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mayron1806/go-clover-core"
)

const (
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)

func GzipMiddleware(level int) clover.Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			// Verifica se o cliente suporta Gzip
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				w.Header().Set("Content-Encoding", "gzip")
				gz, err := gzip.NewWriterLevel(w, level)
				if err != nil {
					http.Error(w, "Failed to apply compression", http.StatusInternalServerError)
					return
				}
				defer gz.Close()

				gzrw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
				next(gzrw, r, ps)
			} else {
				next(w, r, ps)
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
