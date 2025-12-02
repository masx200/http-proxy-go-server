## 1. Implementation

- [x] 1.1 Analyze current upstream handling in main.go
- [x] 1.2 Create OpenSpec change proposal for upstream IP resolution
- [x] 1.3 Design IP resolution and connection fallback mechanism
- [x] 1.4 Add upstreamResolveIPs command-line parameter to main.go
- [x] 1.5 Extend config/types.go with upstream IP resolution configuration
- [x] 1.6 Implement upstream domain to IP resolution function in
      options/options.go
- [x] 1.7 Modify Proxy_net_dial function to support IP-based upstream
      connections
- [x] 1.8 Add sequential IP connection attempt logic with fallback
- [x] 1.9 Add comprehensive logging for connection attempts
- [x] 1.10 Update WebSocket and SOCKS5 dial functions for IP resolution support
- [x] 1.11 Add unit tests for new IP resolution functionality
- [x] 1.12 Add integration tests with simulated DNS pollution scenarios
- [x] 1.13 Update documentation with usage examples
- [x] 1.14 Test backward compatibility with existing configurations
