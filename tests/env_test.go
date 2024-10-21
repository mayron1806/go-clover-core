package tests

import (
	"os"
	"testing"
	"time"

	"github.com/mayron1806/go-clover-core/config"
	"github.com/stretchr/testify/assert"
)

// Mock struct for testing
type MockConfig struct {
	StringField   string        `env:"STRING_FIELD" default:"defaultString"`
	IntField      int           `env:"INT_FIELD" default:"42"`
	BoolField     bool          `env:"BOOL_FIELD" default:"false"`
	DurationField time.Duration `env:"DURATION_FIELD" default:"15m"`
}

func TestLoadEnv_WithDefaults(t *testing.T) {
	// Create a new environment loader
	envLoader := config.NewEnvLoader[MockConfig]()

	// Load environment variables into the struct
	config, err := envLoader.LoadEnv()
	assert.NoError(t, err, "LoadEnv should not return an error")

	// Check if the fields have default values
	assert.Equal(t, "defaultString", config.StringField, "Default value for StringField should be 'defaultString'")
	assert.Equal(t, 42, config.IntField, "Default value for IntField should be 42")
	assert.Equal(t, false, config.BoolField, "Default value for BoolField should be false")
	assert.Equal(t, 15*time.Minute, config.DurationField, "Default value for DurationField should be 15m")
}

func TestLoadEnv_WithEnvVariables(t *testing.T) {
	// Set mock environment variables
	os.Setenv("STRING_FIELD", "customString")
	os.Setenv("INT_FIELD", "100")
	os.Setenv("BOOL_FIELD", "true")
	os.Setenv("DURATION_FIELD", "30m")

	// Clean up environment variables after the test
	defer os.Unsetenv("STRING_FIELD")
	defer os.Unsetenv("INT_FIELD")
	defer os.Unsetenv("BOOL_FIELD")
	defer os.Unsetenv("DURATION_FIELD")

	// Create a new environment loader
	envLoader := config.NewEnvLoader[MockConfig]()

	// Load environment variables into the struct
	config, err := envLoader.LoadEnv()
	assert.NoError(t, err, "LoadEnv should not return an error")

	// Check if the fields have been overridden by environment variables
	assert.Equal(t, "customString", config.StringField, "StringField should be overridden by environment variable")
	assert.Equal(t, 100, config.IntField, "IntField should be overridden by environment variable")
	assert.Equal(t, true, config.BoolField, "BoolField should be overridden by environment variable")
	assert.Equal(t, 30*time.Minute, config.DurationField, "DurationField should be overridden by environment variable")
}

func TestLoadEnv_ValidationFailure(t *testing.T) {
	// Create a mock struct with validation rules
	type InvalidConfig struct {
		MandatoryField string `env:"MANDATORY_FIELD" validate:"required"`
	}

	// Set a mock environment variable without setting MANDATORY_FIELD
	os.Setenv("MANDATORY_FIELD", "")

	// Clean up environment variables after the test
	defer os.Unsetenv("MANDATORY_FIELD")

	// Create a new environment loader
	envLoader := config.NewEnvLoader[InvalidConfig]()

	// Load environment variables into the struct
	_, err := envLoader.LoadEnv()

	// Expecting an error because the mandatory field is not set
	assert.Error(t, err, "LoadEnv should return an error when validation fails")
}

func TestLoadEnv_ValidStruct(t *testing.T) {
	// Create a valid struct to test validation
	type ValidConfig struct {
		MandatoryField string `env:"MANDATORY_FIELD" validate:"required"`
	}

	// Set the mandatory environment variable
	os.Setenv("MANDATORY_FIELD", "Valid value")

	// Clean up environment variables after the test
	defer os.Unsetenv("MANDATORY_FIELD")

	// Create a new environment loader
	envLoader := config.NewEnvLoader[ValidConfig]()

	// Load environment variables into the struct
	config, err := envLoader.LoadEnv()

	// Expecting no error because the mandatory field is set
	assert.NoError(t, err, "LoadEnv should not return an error when validation passes")
	assert.Equal(t, "Valid value", config.MandatoryField, "MandatoryField should be set correctly")
}
