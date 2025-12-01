## 1. Implementation

- [x] 1.1 Analyze current upstream handling in main.go
- [x] 1.2 Create OpenSpec change proposal for upstream IP resolution
- [x] 1.3 Design IP resolution and connection fallback mechanism
- [ ] 1.4 Add upstreamResolveIPs command-line parameter to main.go
- [ ] 1.5 Extend config/types.go with upstream IP resolution configuration
- [ ] 1.6 Implement upstream domain to IP resolution function in options/options.go
- [ ] 1.7 Modify Proxy_net_dial function to support IP-based upstream connections
- [ ] 1.8 Add sequential IP connection attempt logic with fallback
- [ ] 1.9 Add comprehensive logging for connection attempts
- [ ] 1.10 Update WebSocket and SOCKS5 dial functions for IP resolution support
- [ ] 1.11 Add unit tests for new IP resolution functionality
- [ ] 1.12 Add integration tests with simulated DNS pollution scenarios
- [ ] 1.13 Update documentation with usage examples
- [ ] 1.14 Test backward compatibility with existing configurations