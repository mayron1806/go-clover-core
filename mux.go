package clover

import (
	"fmt"
	"net/http"
	"strings"
)

type ServeClover struct {
	handlers map[string]map[string]HandlerFunc
}

func NewServeClover() *ServeClover {
	return &ServeClover{
		handlers: make(map[string]map[string]HandlerFunc),
	}
}

func (mux *ServeClover) Handle(method, pattern string, handler HandlerFunc) {
	if mux.handlers[pattern] == nil {
		mux.handlers[pattern] = make(map[string]HandlerFunc)
	}
	mux.handlers[pattern][method] = handler
}
func (mux *ServeClover) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP", r.URL.Path)

	// iterate over static handlers
	for pattern := range mux.handlers {
		if !isDynamicRoute(pattern) {
			if handlersByMethod, ok := mux.handlers[r.URL.Path]; ok {
				if handler, ok := handlersByMethod[r.Method]; ok {
					handler(newContext(w, r, nil))
					return
				}
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
		}
	}

	// iterate over dynamic handlers
	for pattern, handlersByMethod := range mux.handlers {
		// Verificar se o padrão da rota contém um parâmetro dinâmico
		if isDynamicRoute(pattern) {
			if params, match := extractParams(pattern, r.URL.Path); match {
				if handler, ok := handlersByMethod[r.Method]; ok {
					fmt.Println("Dynamic route", r.URL.Path)
					// Cria o contexto com os parâmetros capturados
					handler(newContext(w, r, params))
					return
				}
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
		} else if pattern == r.URL.Path {
			// Caso o padrão seja exato
			if handler, ok := handlersByMethod[r.Method]; ok {
				handler(newContext(w, r, nil))
				return
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	fmt.Println("Not found")
	http.NotFound(w, r)
}

// Verifica se a rota contém parâmetros dinâmicos, ex: "/users/{id}"
func isDynamicRoute(pattern string) bool {
	return strings.Contains(pattern, "{") && strings.Contains(pattern, "}")
}

// Função para extrair parâmetros dinâmicos de uma rota
func extractParams(pattern, path string) (map[string]string, bool) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	// Verifica se o tamanho do caminho e do padrão são compatíveis
	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(map[string]string)
	for i := range patternParts {
		if strings.HasPrefix(patternParts[i], "{") && strings.HasSuffix(patternParts[i], "}") {
			paramName := strings.TrimSuffix(strings.TrimPrefix(patternParts[i], "{"), "}")
			params[paramName] = pathParts[i]
		} else if patternParts[i] != pathParts[i] {
			return nil, false
		}
	}
	return params, true
}
