package jwt_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mayron1806/go-clover-core"
	"github.com/mayron1806/go-clover-core/jwt"
	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware_HeaderTokenSuccess(t *testing.T) {
	// Configura o JWT middleware
	opts, err := jwt.DefaultJWTMiddlewareOptions()
	assert.NoError(t, err)

	token, err := jwt.GenerateToken(jwt.Subject{"test-subject": "teste"}, opts.JWTOpts)
	assert.NoError(t, err)

	middleware := jwt.JWTMiddleware(opts)

	// Cria uma requisição simulada com o token no cabeçalho Authorization
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	// Contexto simulado do Clover
	ctx := clover.NewContext(w, req, map[string]string{})

	// Executa o middleware
	nextCalled := false
	next := func(c *clover.Context) {
		nextCalled = true
	}

	middleware(next)(ctx)

	// Verifica se o middleware permitiu a execução da próxima função
	assert.True(t, nextCalled)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestJWTMiddleware_MissingToken(t *testing.T) {
	// Configura o JWT middleware
	opts, err := jwt.DefaultJWTMiddlewareOptions()
	assert.NoError(t, err)

	middleware := jwt.JWTMiddleware(opts)

	// Cria uma requisição simulada sem o token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// Contexto simulado do Clover
	ctx := clover.NewContext(w, req, map[string]string{})

	// Executa o middleware
	next := func(c *clover.Context) {}

	middleware(next)(ctx)

	// Verifica se o middleware retornou status 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestJWTMiddleware_InvalidTokenFormat(t *testing.T) {
	// Configura o JWT middleware
	opts, err := jwt.DefaultJWTMiddlewareOptions()
	assert.NoError(t, err)

	middleware := jwt.JWTMiddleware(opts)

	// Cria uma requisição simulada com um formato de token inválido
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "InvalidToken")
	w := httptest.NewRecorder()

	// Contexto simulado do Clover
	ctx := clover.NewContext(w, req, map[string]string{})

	// Executa o middleware
	next := func(c *clover.Context) {}

	middleware(next)(ctx)

	// Verifica se o middleware retornou status 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
