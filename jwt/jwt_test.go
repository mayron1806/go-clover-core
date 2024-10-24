package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Mock TokenOptions for testing
func mockTokenOptions() JWTTokenOptions {
	return JWTTokenOptions{
		Duration: time.Minute * 15,
		Issuer:   "test-issuer",
		Secret:   "test-secret",
		Method:   "HS256",
	}
}

func TestGenerateToken(t *testing.T) {
	// Mock options
	opts := mockTokenOptions()

	// Payload to include in the token
	payload := "test-subject"

	// Generate token
	token, err := GenerateToken(Subject{}, opts)
	assert.NoError(t, err, "Error generating token")
	fmt.Println(token)
	assert.NotEmpty(t, token, "Generated token should not be empty")

	// Parse the token to validate
	parsedSubject, err := ParseToken(token, opts)
	assert.NoError(t, err, "Error parsing token")
	assert.Equal(t, payload, parsedSubject, "The parsed subject should match the payload")
}

func TestParseToken_ValidToken(t *testing.T) {
	// Mock options
	opts := mockTokenOptions()

	// Generate a valid token
	payload := "test-subject"
	token, err := GenerateToken(Subject{}, opts)
	assert.NoError(t, err, "Error generating token")

	// Parse the token
	parsedSubject, err := ParseToken(token, opts)
	assert.NoError(t, err, "Error parsing token")
	assert.Equal(t, payload, parsedSubject, "The parsed subject should match the payload")
}

func TestParseToken_InvalidToken(t *testing.T) {
	// Initialize the JWT service

	// Mock options
	opts := mockTokenOptions()

	// Pass an invalid token
	invalidToken := "invalid.token.string"

	// Parse the token
	parsedSubject, err := ParseToken(invalidToken, opts)
	assert.Error(t, err, "Expected an error for invalid token")
	assert.Empty(t, parsedSubject, "Parsed subject should be empty for invalid token")
}

func TestDefaultTokenOptions(t *testing.T) {
	// Set expected environment variables for the test
	opts, err := DefaultTokenOptions()

	assert.NoError(t, err, "Error loading default token options")
	assert.NotEmpty(t, opts.Secret, "Token secret should be set")
	assert.Equal(t, time.Minute*15, opts.Duration, "Token duration should be 15 minutes")
	assert.Equal(t, "HS256", opts.Method, "Default method should be HS256")
}

func TestGenerateToken_Expiration(t *testing.T) {

	// Mock options with a short expiration
	opts := JWTTokenOptions{
		Duration: time.Second * 1, // 1 second for quick expiration
		Issuer:   "test-issuer",
		Secret:   "test-secret",
		Method:   "HS256",
	}

	// Generate token
	token, err := GenerateToken(Subject{}, opts)
	assert.NoError(t, err, "Error generating token")

	// Sleep for 2 seconds to let the token expire
	time.Sleep(2 * time.Second)

	// Parse the expired token
	parsedSubject, err := ParseToken(token, opts)
	assert.Error(t, err, "Expected an error for expired token")
	assert.Empty(t, parsedSubject, "Parsed subject should be empty for expired token")
}