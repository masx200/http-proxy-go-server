# Tests Directory

This directory contains test files and configurations for the HTTP proxy Go
server.

## Structure

- `config/` - Configuration files for testing
- `standalone/` - Standalone test applications with main functions
- `ProcessManager.go` - Test process management utilities

## Configuration Files

### Test Configurations (`config/`)

- `config.simple.json` - Simple JSON configuration with basic settings
- `config.simple.yaml` - Simple YAML configuration with basic settings
- `config.example.json` - Complete JSON configuration example with all features
- `config.example.yaml` - Complete YAML configuration example with all features
- `config.invalid.json` - Invalid configuration for testing validation error
  handling

### Usage Examples

```bash
# Test with JSON configuration
go run ./cmd/ -config tests/config/config.simple.json

# Test with YAML configuration
go run ./cmd/ -config tests/config/config.simple.yaml

# Test with complex configuration
go run ./cmd/ -config tests/config/config.example.json

# Test validation with invalid configuration (should show errors)
go run ./cmd/ -config tests/config/config.invalid.json
```

## Standalone Tests (`standalone/`)

Each standalone test is in its own directory to avoid package conflicts:

- `test_aof_simple/` - Simple AOF (Append Only File) test
- `test_dnscache_aof/` - DNS cache with AOF test
- `test_fixed/` - Fixed functionality test
- `test_final/` - Final integration test
- `test_performance/` - Performance test

### Running Standalone Tests

```bash
# Run specific standalone tests
cd tests/standalone/test_aof_simple && go run .
cd tests/standalone/test_performance && go run .
```

## Configuration Validation

The project includes comprehensive JSON Schema validation for configuration
files:

- Validates types, formats, and value ranges
- Enforces conditional requirements (e.g., TLS cert requires key)
- Supports both JSON and YAML formats
- Provides clear error messages for invalid configurations

### Schema Location

The JSON Schema is located at `config/config.schema.json` and is embedded in the
binary during build.
