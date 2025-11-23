package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"embed"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed config.schema.json
var schemaFS embed.FS

// ValidationError represents a configuration validation error
type ValidationError struct {
	Message   string
	Errors    []string
	FieldName string
}

func (e *ValidationError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("%s:\n%s", e.Message, strings.Join(e.Errors, "\n"))
	}
	return e.Message
}

// LoadAndValidateConfig loads configuration from a file and validates it against the JSON schema
func LoadAndValidateConfig(configPath string) (*Config, error) {
	// Load and validate the configuration
	config, err := loadConfigFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return config, nil
}

// loadConfigFile loads configuration from JSON or YAML file with validation
func loadConfigFile(configPath string) (*Config, error) {
	// Read the configuration file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Convert to JSON based on file extension
	ext := strings.ToLower(filepath.Ext(configPath))
	var jsonData []byte

	switch ext {
	case ".json":
		jsonData = data
	case ".yaml", ".yml":
		// Convert YAML to JSON
		var yamlData interface{}
		if err := yaml.Unmarshal(data, &yamlData); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %w", err)
		}

		jsonData, err = json.Marshal(yamlData)
		if err != nil {
			return nil, fmt.Errorf("failed to convert YAML to JSON: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported config format: %s (supported: .json, .yaml, .yml)", ext)
	}

	// Validate against JSON schema
	if err := validateJSONSchema(jsonData); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	// Parse into Config struct
	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}

	// Apply default values
	applyDefaults(&config)

	return &config, nil
}

// validateJSONSchema validates JSON data against the embedded schema
func validateJSONSchema(jsonData []byte) error {
	// Load schema from embedded filesystem
	schemaBytes, err := fs.ReadFile(schemaFS, "config.schema.json")
	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(schemaBytes)
	documentLoader := gojsonschema.NewBytesLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("schema validation error: %w", err)
	}

	if !result.Valid() {
		var errors []string
		for _, desc := range result.Errors() {
			errors = append(errors, desc.String())
		}
		return &ValidationError{
			Message: "Configuration validation failed",
			Errors:  errors,
		}
	}

	return nil
}

// applyDefaults applies default values to the configuration
func applyDefaults(config *Config) {
	// Apply server defaults
	if config.Hostname == "" {
		config.Hostname = "0.0.0.0"
	}
	if config.Port == 0 {
		config.Port = 8080
	}

	// Apply DNS cache defaults
	if !config.DNSCache.EnabledSet {
		config.DNSCache.Enabled = true
	}
	if config.DNSCache.File == "" {
		config.DNSCache.File = "./dns_cache.json"
	}
	if config.DNSCache.TTL == "" {
		config.DNSCache.TTL = "10m"
	}
	if config.DNSCache.SaveInterval == "" {
		config.DNSCache.SaveInterval = "30s"
	}
	if !config.DNSCache.AOFEnabledSet {
		config.DNSCache.AOFEnabled = true
	}
	if config.DNSCache.AOFFile == "" {
		config.DNSCache.AOFFile = "./dns_cache.aof"
	}
	if config.DNSCache.AOFInterval == "" {
		config.DNSCache.AOFInterval = "1s"
	}

	// Validate and parse duration strings
	if err := validateAndParseDurations(config); err != nil {
		// Log warning but don't fail - let the application handle it
		fmt.Printf("Warning: %v\n", err)
	}
}

// validateAndParseDurations validates duration strings and parses them
func validateAndParseDurations(config *Config) error {
	// Validate TTL
	if _, err := time.ParseDuration(config.DNSCache.TTL); err != nil {
		return fmt.Errorf("invalid DNS cache TTL '%s': %w", config.DNSCache.TTL, err)
	}

	// Validate SaveInterval
	if _, err := time.ParseDuration(config.DNSCache.SaveInterval); err != nil {
		return fmt.Errorf("invalid DNS cache save interval '%s': %w", config.DNSCache.SaveInterval, err)
	}

	// Validate AOFInterval
	if _, err := time.ParseDuration(config.DNSCache.AOFInterval); err != nil {
		return fmt.Errorf("invalid DNS cache AOF interval '%s': %w", config.DNSCache.AOFInterval, err)
	}

	return nil
}

// ValidateCommandLineConfig validates configuration built from command line arguments
func ValidateCommandLineConfig(config *Config) error {
	// Convert config to JSON for validation
	jsonData, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to convert config to JSON: %w", err)
	}

	// Validate against schema
	if err := validateJSONSchema(jsonData); err != nil {
		return fmt.Errorf("command line configuration validation failed: %w", err)
	}

	// Apply defaults
	applyDefaults(config)

	return nil
}

// GetSchema returns the JSON schema as a string (useful for documentation)
func GetSchema() (string, error) {
	schemaBytes, err := fs.ReadFile(schemaFS, "config.schema.json")
	if err != nil {
		return "", fmt.Errorf("failed to read schema: %w", err)
	}
	return string(schemaBytes), nil
}
