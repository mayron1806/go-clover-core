package config

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/mayron1806/go-fast/logging"
)

type DefaultEnvLoader[T any] struct {
	logger *logging.Logger
}

// Função para criar um novo DefaultEnvLoader
func NewEnvLoader[T any]() *DefaultEnvLoader[T] {
	return &DefaultEnvLoader[T]{
		logger: logging.NewLogger("Env Loader"),
	}
}

// Método LoadEnv para carregar as variáveis de ambiente em uma struct
func (c *DefaultEnvLoader[T]) LoadEnv() (*T, error) {
	if envData, success := env.(T); success {
		return &envData, nil
	}
	c.logger.Info("attempting to load .env file from project root...")

	// Obtém o caminho do diretório onde o go.mod está localizado (project root)
	projectRoot, err := getProjectRoot()
	if err != nil {
		c.logger.Warnf("project root not found: %s. Falling back to system environment variables", err.Error())
	}

	// Tenta carregar o arquivo .env da raiz do projeto
	if projectRoot != "" {
		envPath := filepath.Join(projectRoot, ".env")
		err = godotenv.Load(envPath)
		if err != nil {
			c.logger.Warnf("no .env file found at %s, continuing with system environment variables", envPath)
		} else {
			c.logger.Infof(".env file loaded from %s", envPath)
		}
	}

	// Cria uma instância da struct
	env := new(T)

	// Carrega as variáveis de ambiente na struct
	if err := loadEnvIntoStruct(env); err != nil {
		c.logger.Errorf("error loading environment variables into struct: %s", err.Error())
		return nil, err
	}

	// Valida os campos da struct
	envValidator := validator.New()
	err = envValidator.Struct(env)
	if err != nil {
		c.logger.Errorf("error validating struct: %s", err.Error())
		return nil, err
	}

	c.logger.Info("environment variables loaded and validated successfully")
	return env, nil
}

// Função para obter o diretório onde o arquivo go.mod está localizado
func getProjectRoot() (string, error) {
	// Obtém o diretório atual
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Procura pelo arquivo go.mod como referência do diretório raiz do projeto
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Se chegar à raiz do sistema e não encontrar, retorna erro
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("could not find project root (go.mod) file")
		}

		// Sobe um nível no diretório
		dir = parent
	}
}

// Função para carregar os valores do .env na struct
func loadEnvIntoStruct(config interface{}) error {
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		envKey := fieldType.Tag.Get("env")
		defaultValue := fieldType.Tag.Get("default")

		if envKey == "" {
			envKey = fieldType.Name
		}

		envValue := getEnv(envKey, defaultValue)

		if field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(envValue)
			case reflect.Int:
				intValue, _ := strconv.Atoi(envValue)
				field.SetInt(int64(intValue))
			case reflect.Bool:
				boolValue := envValue == "true"
				field.SetBool(boolValue)
			case reflect.Int64:
				if field.Type() == reflect.TypeOf(time.Duration(0)) {
					duration, _ := time.ParseDuration(envValue)
					field.SetInt(int64(duration))
				}
			}
		}
	}

	return nil
}

// Função para obter o valor de uma variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
