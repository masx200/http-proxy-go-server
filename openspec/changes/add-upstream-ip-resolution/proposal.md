## Why

The current upstream proxy connection system sends domain names directly to upstream proxies, which can fail when DNS pollution or censorship affects domain resolution for the upstream proxy itself. By resolving upstream proxy domains to IP addresses before connection and attempting each IP sequentially, we can bypass DNS-based blocking and ensure reliable upstream connectivity in restricted network environments.

## What Changes

- Add new command-line parameter `-upstream-resolve-ips` to enable IP-based upstream connections
- Modify upstream connection logic to resolve domain names to IP addresses using configured DNS/DoH services
- Implement sequential IP connection fallback mechanism with detailed logging
- Add connection retry logic that tries each resolved IP until successful connection
- Maintain backward compatibility with existing domain-based upstream connections

## Impact

- **Affected specs**: Upstream proxy connections, DNS resolution
- **Affected code**: `cmd/main.go`, `options/options.go`, upstream connection logic in transport configurations
- **Breaking changes**: None (feature is opt-in via new command-line parameter)
- **Performance**: Improved connection reliability in censored environments, slightly increased initial connection setup time