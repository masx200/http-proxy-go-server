## Context

The current upstream proxy system connects to upstream proxies using domain
names directly. In environments with DNS pollution or censorship, this approach
can fail when the upstream proxy's domain name is blocked or resolves to
incorrect addresses. The project already has a robust DNS/DoH infrastructure
that can reliably resolve domains using multiple DNS over HTTPS services.

## Goals / Non-Goals

- **Goals**:
  - Enable reliable upstream proxy connections in censored environments
  - Leverage existing DNS/DoH infrastructure for upstream domain resolution
  - Maintain full backward compatibility with existing configurations
  - Provide detailed logging for debugging connection issues
  - Support all existing upstream proxy types (HTTP, SOCKS5, WebSocket)

- **Non-Goals**:
  - Modify existing DNS/DoH infrastructure (reuse as-is)
  - Change authentication mechanisms for upstream proxies
  - Alter configuration file format
  - Impact performance when feature is disabled

## Decisions

- **Decision**: Add new command-line parameter `-upstream-resolve-ips` to enable
  IP-based upstream connections
- **Rationale**: Command-line approach maintains backward compatibility and
  doesn't require configuration file changes. Opt-in behavior ensures existing
  deployments continue working without modification.

- **Decision**: Implement connection fallback mechanism that tries each resolved
  IP sequentially
- **Rationale**: Simple, reliable approach that ensures at least one working IP
  will be used if DNS resolution succeeds

- **Decision**: Use existing DNS/DoH infrastructure for upstream resolution
- **Rationale**: Leverages proven, tested DNS resolution code with multiple DoH
  providers and caching

- **Decision**: Add detailed logging for connection attempts
- **Rationale**: Essential for debugging connection issues in censored
  environments

## Risks / Trade-offs

- **Risk**: Increased connection setup time when resolving upstream domains
  - **Mitigation**: DNS caching in existing infrastructure reduces resolution
    overhead

- **Risk**: Potential connection loops if resolved IPs point to the same failing
  infrastructure
  - **Mitigation**: Random IP ordering and connection timeout configuration

- **Trade-off**: Slight complexity increase in upstream connection logic
  - **Mitigation**: Clear separation of IP-based and domain-based connection
    paths

## Migration Plan

1. Add new command-line parameter parsing to main.go
2. Extend UpStream configuration with IP resolution metadata
3. Modify Proxy_net_dial function to support IP-based upstream connections
4. Add upstream IP resolution utility functions
5. Implement sequential IP connection attempt logic
6. Add comprehensive logging for connection attempts
7. Update documentation and examples

## Open Questions

- Should the IP resolution timeout be configurable separately from general DNS
  timeouts?
- Should successful IP connections be cached for faster subsequent connections?
- Should we support mixed configurations (some upstreams use IP resolution,
  others don't)?
