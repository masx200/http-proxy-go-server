# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

`http-proxy-go-server` is a feature-rich HTTP/HTTPS proxy server implemented in
Go that supports multiple upstream proxy types, DNS over HTTPS (DoH), and
flexible routing configurations. The server acts as a forward proxy that can
route requests through various upstream proxies based on configurable rules and
patterns.

## Architecture

### Core Components

#### 1. Main Entry Point (`cmd/main.go`)

- **Purpose**: Central application entry point with configuration management
- **Key Features**:
  - Command-line argument parsing with extensive flag support
  - JSON configuration file loading and merging
  - Dynamic upstream proxy configuration
  - Signal handling for graceful shutdown
  - Proxy selection logic with CIDR matching

#### 2. Proxy Server Modes

The server operates in four different modes based on configuration:

1. **Simple Mode** (`simple/`): Basic HTTP proxy without authentication
2. **Auth Mode** (`auth/`): HTTP proxy with Basic authentication
3. **TLS Mode** (`tls/`): HTTPS proxy with TLS encryption
4. **TLS+Auth Mode** (`tls+auth/`): HTTPS proxy with both TLS and authentication

#### 3. HTTP Server Component (`http/`)

- **Purpose**: Internal HTTP server for handling HTTP requests
- **Features**:
  - Gin-based HTTP server for internal routing
  - Forwarded header parsing and manipulation
  - Request forwarding logic
  - Authentication handling

#### 4. Options Management (`options/`)

- **Purpose**: Network dialing and DoH resolution
- **Key Functions**:
  - `Proxy_net_Dial()`: Custom network dialing with DoH support
  - IP address shuffling for load balancing
  - Multiple DoH server support with fallback
  - Local hosts file integration

#### 5. Connection Handling (`connect/`)

- **Purpose**: Handle HTTP CONNECT method for HTTPS tunneling
- **Features**:
  - CONNECT method parsing
  - Upstream proxy connection establishment
  - TCP tunneling for HTTPS traffic

#### 6. DNS Resolution (`doh/`, `hosts/`)

- **DoH Integration**: DNS over HTTPS support with HTTP/2 and HTTP/3
- **Local Hosts**: Custom hosts file resolution
- **Resolver Interface**: Abstracted DNS resolution system

#### 7. Transport Configuration (`resolver/`)

- **Purpose**: Defines interfaces for name resolution
- **Abstraction**: Provides pluggable DNS resolver implementations

#### 8. Utilities (`utils/`)

- **Purpose**: Common utility functions
- **Functions**: Proxy URL processing and validation

## Configuration System

### Configuration Structure

```go
type Config struct {
    Hostname   string      `json:"hostname"`
    Port       int         `json:"port"`
    ServerCert string      `json:"server_cert"`
    ServerKey  string      `json:"server_key"`
    Username   string      `json:"username"`
    Password   string      `json:"password"`
    Doh        []DohConfig `json:"doh"`
    UpStreams  map[string]UpStream `json:"upstreams"`
    Rules      []struct {
        Filter   string `json:"filter"`
        Upstream string `json:"upstream"`
    } `json:"rules"`
    Filters    map[string]struct {
        Patterns []string `json:"patterns"`
    } `json:"filters"`
}
```

### Upstream Proxy Types

The system supports multiple upstream proxy types:

```go
type UpStream struct {
    TYPE        string   `json:"type"`
    HTTP_PROXY  string   `json:"http_proxy"`
    HTTPS_PROXY string   `json:"https_proxy"`
    BypassList  []string `json:"bypass_list"`

    // WebSocket support
    WS_PROXY    string `json:"ws_proxy"`
    WS_USERNAME string `json:"ws_username"`
    WS_PASSWORD string `json:"ws_password"`

    // SOCKS5 support
    SOCKS5_PROXY    string `json:"socks5_proxy"`
    SOCKS5_USERNAME string `json:"socks5_username"`
    SOCKS5_PASSWORD string `json:"socks5_password"`

    // HTTP authentication
    HTTP_USERNAME string `json:"http_username"`
    HTTP_PASSWORD string `json:"http_password"`
}
```

## Key Features

### 1. Multi-Protocol Support

- **HTTP/HTTPS Proxy**: Standard HTTP proxy functionality
- **WebSocket Proxy**: Forward requests through WebSocket upstreams
- **SOCKS5 Proxy**: SOCKS5 protocol support with authentication
- **DoH Integration**: DNS over HTTPS for encrypted DNS resolution

### 2. Advanced Routing

- **Pattern Matching**: Wildcard and exact domain matching
- **CIDR Support**: IP range-based routing decisions
- **Bypass Lists**: Domain/IP exclusion lists
- **Filter System**: Flexible rule-based routing

### 3. Authentication & Security

- **Basic Auth**: HTTP Basic authentication support
- **TLS Encryption**: HTTPS proxy with custom certificates
- **Upstream Auth**: Authentication for upstream proxies
- **Credential Override**: Configuration-based credential management

### 4. DoH (DNS over HTTPS)

- **Multiple DoH Servers**: Configurable DoH endpoints
- **HTTP/2 & HTTP/3**: Support for modern ALPN protocols
- **Fallback Mechanism**: Multiple DNS resolution strategies
- **Local DNS Integration**: Hosts file resolution

### 5. Performance Features

- **IP Shuffling**: Random IP selection for load balancing
- **Connection Pooling**: Efficient connection reuse
- **Timeout Management**: Configurable timeouts
- **Graceful Shutdown**: Clean process termination

## Build & Development

### Build Commands

```bash
# Build the main executable
go build -o main.exe ./cmd/main.go

# Build with optimizations
go build -ldflags="-s -w" -o main.exe ./cmd/main.go

# Run directly
go run ./cmd/main.go [options]
```

### Docker Support

Multi-stage Dockerfile with:

- Go 1.24.4 build environment
- Optimized Alpine runtime image
- Dependency caching
- Binary optimization

### Dependencies

Key external dependencies:

- `github.com/gin-gonic/gin`: HTTP web framework
- `github.com/masx200/socks5-websocket-proxy-golang`: SOCKS5/WebSocket client
- `github.com/tantalor93/doq-go`: DoH client
- `github.com/miekg/dns`: DNS library
- `github.com/gorilla/websocket`: WebSocket library

## Configuration Examples

### Basic HTTP Proxy

```json
{
  "hostname": "0.0.0.0",
  "port": 8080,
  "username": "admin",
  "password": "secret",
  "doh": [
    {
      "ip": "8.8.8.8",
      "alpn": "h2",
      "url": "https://dns.google/dns-query"
    }
  ]
}
```

### Upstream Proxy Configuration

```json
{
  "upstreams": {
    "ws_proxy": {
      "type": "websocket",
      "ws_proxy": "ws://127.0.0.1:1081",
      "ws_username": "user",
      "ws_password": "pass"
    },
    "socks5_proxy": {
      "type": "socks5",
      "socks5_proxy": "socks5://127.0.0.1:1080",
      "socks5_username": "user",
      "socks5_password": "pass"
    },
    "http_proxy": {
      "type": "http",
      "http_proxy": "http://127.0.0.1:8080",
      "https_proxy": "http://127.0.0.1:8080",
      "http_username": "user",
      "http_password": "pass"
    }
  },
  "rules": [
    {
      "filter": "ws_filter",
      "upstream": "ws_proxy"
    },
    {
      "filter": "socks5_filter",
      "upstream": "socks5_proxy"
    },
    {
      "filter": "http_filter",
      "upstream": "http_proxy"
    }
  ],
  "filters": {
    "ws_filter": {
      "patterns": ["*.example.com"]
    },
    "socks5_filter": {
      "patterns": ["*.internal"]
    },
    "http_filter": {
      "patterns": ["*"]
    }
  }
}
```

### Command Line Usage

```bash
# Basic proxy
go run ./cmd/main.go -hostname 0.0.0.0 -port 8080

# With authentication
go run ./cmd/main.go -username admin -password secret

# With TLS
go run ./cmd/main.go -server_cert cert.pem -server_key key.pem

# With upstream proxy
go run ./cmd/main.go -upstream-type websocket -upstream-address ws://127.0.0.1:1081

# With DoH
go run ./cmd/main.go -dohurl https://dns.google/dns-query -dohip 8.8.8.8 -dohalpn h2

# From config file
go run ./cmd/main.go -config config.json
```

## Testing

### Test Structure

- **Integration Tests** (`tests/`): End-to-end proxy functionality testing
- **Process Management**: Automated process lifecycle handling
- **HTTP/HTTPS Testing**: curl-based functional testing
- **Timeout Handling**: Robust test timeout mechanisms

### Running Tests

```bash
# Run proxy tests
cd tests && go test -v

# Manual testing with curl
curl -x http://localhost:8080 http://example.com
curl -x http://localhost:8080 https://example.com
```

## Development Workflow

### 1. Code Organization

```
├── cmd/                    # Application entry point
├── simple/                 # Simple proxy mode
├── auth/                   # Authenticated proxy mode
├── tls/                    # TLS proxy mode
├── tls+auth/               # TLS+Auth proxy mode
├── http/                   # HTTP server components
├── options/                # Network and DNS options
├── connect/                # CONNECT method handling
├── doh/                    # DNS over HTTPS
├── hosts/                  # Local hosts resolution
├── resolver/               # DNS resolver interfaces
├── utils/                  # Utility functions
└── tests/                  # Integration tests
```

### 2. Development Commands

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run tests
go test ./...

# Build for production
go build -ldflags="-s -w" -o http-proxy-go-server ./cmd/main.go

# Build Docker image
docker build -t http-proxy-go-server .
```

## Security Considerations

1. **Authentication**: Always use strong passwords for proxy authentication
2. **TLS**: Use valid certificates for HTTPS proxy mode
3. **Upstream Security**: Ensure upstream proxy credentials are secure
4. **Network Isolation**: Consider network access controls
5. **Logging**: Monitor proxy access logs for suspicious activity

## Performance Optimization

1. **Connection Pooling**: Configure appropriate keep-alive settings
2. **DoH Caching**: Leverage DoH response caching
3. **IP Shuffling**: Use multiple IP addresses for load balancing
4. **Timeout Configuration**: Set appropriate timeouts for your use case
5. **Resource Limits**: Configure connection limits and resource constraints

## Troubleshooting

### Common Issues

1. **Port Conflicts**: Ensure ports are not already in use
2. **Certificate Issues**: Verify TLS certificate and key files
3. **Network Connectivity**: Check upstream proxy accessibility
4. **DNS Resolution**: Verify DoH server availability
5. **Authentication**: Confirm credential configuration

### Debug Commands

```bash
# Check port usage
netstat -ano | findstr :8080

# Test proxy connectivity
curl -v -x http://localhost:8080 http://example.com

# Check DNS resolution
nslookup example.com
```

## Contributing

1. **Code Style**: Follow Go formatting standards
2. **Testing**: Add tests for new features
3. **Documentation**: Update documentation for changes
4. **Security**: Consider security implications of changes
5. **Performance**: Test performance impact of modifications

This comprehensive architecture provides a solid foundation for understanding
and extending the http-proxy-go-server codebase. The modular design allows for
easy customization and enhancement while maintaining stability and performance.
