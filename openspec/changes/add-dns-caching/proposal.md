## Why
The http-proxy-go-server currently performs DNS resolution for each request without caching, resulting in redundant DNS queries for the same domains. Adding DNS caching will improve performance, reduce external DNS service load, and decrease response latency for repeated domain resolutions.

## What Changes
- Add DNS caching module with file-based persistence
- Integrate cache into existing NameResolver interface implementations
- Add command-line parameters for cache configuration
- Implement automatic cache expiration (default 10 minutes TTL)
- Support for different DNS record types (A, AAAA, CNAME, etc.)

## Impact
- **Affected specs**: dns-resolution capability
- **Affected code**:
  - `resolver/resolver.go` - Add caching layer to resolver implementations
  - `cmd/main.go` - Add cache-related command-line flags
  - New package `dnscache/` - Cache implementation
  - All proxy modes (simple, auth, tls, tls+auth) will benefit from cached DNS resolution