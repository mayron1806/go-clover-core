package cors

import (
	"net/http"
	"strings"
)

type CORSOptions struct {
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
}

func Cors(opts CORSOptions) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(opts.AllowedOrigins, ","))
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowedMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowedHeaders, ","))

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	})
}
