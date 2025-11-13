## 1. DNS Cache Implementation
- [ ] 1.1 Create `dnscache` package with PatrickMN/go-cache library
- [ ] 1.2 Implement DNSCache struct with file persistence
- [ ] 1.3 Add cache key generation (DNS type + domain normalization)
- [ ] 1.4 Implement atomic save/load operations with temp files
- [ ] 1.5 Add background save ticker and graceful shutdown
- [ ] 1.6 Add cache statistics and monitoring methods

## 2. Command Line Integration
- [ ] 2.1 Add cache-related flags to main.go:
  - `-cache-file` (default: "./dns_cache.json")
  - `-cache-ttl` (default: "10m")
  - `-cache-save-interval` (default: "30s")
  - `-cache-enabled` (default: true)
- [ ] 2.2 Parse duration strings for TTL and interval
- [ ] 2.3 Add cache configuration to Config struct
- [ ] 2.4 Integrate cache flags with existing config file loading

## 3. Resolver Integration
- [ ] 3.1 Create CachingResolver wrapper that implements NameResolver interface
- [ ] 3.2 Modify CreateHostsResolver to use caching layer
- [ ] 3.3 Modify CreateDOHResolver to use caching layer
- [ ] 3.4 Modify CreateDOH3Resolver to use caching layer
- [ ] 3.5 Modify CreateHostsAndDohResolver to use caching layer
- [ ] 3.6 Ensure cache is shared across all resolver types

## 4. Proxy Mode Integration
- [ ] 4.1 Update simple proxy mode to pass cache configuration
- [ ] 4.2 Update auth proxy mode to pass cache configuration
- [ ] 4.3 Update tls proxy mode to pass cache configuration
- [ ] 4.4 Update tls+auth proxy mode to pass cache configuration
- [ ] 4.5 Ensure cache is properly initialized in all modes

## 5. Testing and Validation
- [ ] 5.1 Create unit tests for DNS cache operations
- [ ] 5.2 Test cache persistence across application restarts
- [ ] 5.3 Test cache TTL expiration
- [ ] 5.4 Test concurrent access to cache
- [ ] 5.5 Test integration with all resolver types
- [ ] 5.6 Test command-line parameter parsing
- [ ] 5.7 Test invalid cache file handling

## 6. Dependencies and Build
- [ ] 6.1 Add PatrickMN/go-cache dependency to go.mod
- [ ] 6.2 Update go.sum with new dependency
- [ ] 6.3 Verify build succeeds with all proxy modes
- [ ] 6.4 Test with existing configuration files

## 7. Documentation
- [ ] 7.1 Update CLAUDE.md with cache configuration examples
- [ ] 7.2 Add cache configuration section to resolver_architecture.md
- [ ] 7.3 Update README.md with cache feature description
- [ ] 7.4 Add command-line help text for new flags