package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/mayron1806/go-clover-core/logging"
	"github.com/mayron1806/go-clover-core/service"
	"github.com/stretchr/testify/assert"
)

// Mock TokenOptions for testing
func mockTokenOptions() service.TokenOptions {
	return service.TokenOptions{
		Duration: time.Minute * 15,
		Issuer:   "test-issuer",
		Secret:   "test-secret",
		Method:   "HS256",
	}
}

func TestGenerateToken(t *testing.T) {
	// Initialize the JWT service
	jwtService := service.NewJWTService()

	// Mock options
	opts := mockTokenOptions()

	// Payload to include in the token
	payload := "test-subject"

	// Generate token
	token, err := jwtService.GenerateToken(payload, opts)
	assert.NoError(t, err, "Error generating token")
	fmt.Println(token)
	assert.NotEmpty(t, token, "Generated token should not be empty")

	// Parse the token to validate
	parsedSubject, err := jwtService.ParseToken(token, opts)
	assert.NoError(t, err, "Error parsing token")
	assert.Equal(t, payload, parsedSubject, "The parsed subject should match the payload")
}

func TestParseToken_ValidToken(t *testing.T) {
	// Initialize the JWT service
	jwtService := service.NewJWTService()

	// Mock options
	opts := mockTokenOptions()

	// Generate a valid token
	payload := "test-subject"
	token, err := jwtService.GenerateToken(payload, opts)
	assert.NoError(t, err, "Error generating token")

	// Parse the token
	parsedSubject, err := jwtService.ParseToken(token, opts)
	assert.NoError(t, err, "Error parsing token")
	assert.Equal(t, payload, parsedSubject, "The parsed subject should match the payload")
}

func TestParseToken_InvalidToken(t *testing.T) {
	// Initialize the JWT service
	jwtService := service.NewJWTService()

	// Mock options
	opts := mockTokenOptions()

	// Pass an invalid token
	invalidToken := "invalid.token.string"

	// Parse the token
	parsedSubject, err := jwtService.ParseToken(invalidToken, opts)
	assert.Error(t, err, "Expected an error for invalid token")
	assert.Empty(t, parsedSubject, "Parsed subject should be empty for invalid token")
}

func TestDefaultTokenOptions(t *testing.T) {
	// Set expected environment variables for the test
	logger := logging.NewLogger("TestDefaultTokenOptions")
	opts, err := service.DefaultTokenOptions()
	logger.Debugf("%+v", opts)

	assert.NoError(t, err, "Error loading default token options")
	assert.NotEmpty(t, opts.Secret, "Token secret should be set")
	assert.Equal(t, time.Minute*15, opts.Duration, "Token duration should be 15 minutes")
	assert.Equal(t, "HS256", opts.Method, "Default method should be HS256")
}

func TestGenerateToken_Expiration(t *testing.T) {
	// Initialize the JWT service
	jwtService := service.NewJWTService()

	// Mock options with a short expiration
	opts := service.TokenOptions{
		Duration: time.Second * 1, // 1 second for quick expiration
		Issuer:   "test-issuer",
		Secret:   "test-secret",
		Method:   "HS256",
	}

	// Generate token
	payload := "test-subject"
	token, err := jwtService.GenerateToken(payload, opts)
	assert.NoError(t, err, "Error generating token")

	// Sleep for 2 seconds to let the token expire
	time.Sleep(2 * time.Second)

	// Parse the expired token
	parsedSubject, err := jwtService.ParseToken(token, opts)
	assert.Error(t, err, "Expected an error for expired token")
	assert.Empty(t, parsedSubject, "Parsed subject should be empty for expired token")
}
