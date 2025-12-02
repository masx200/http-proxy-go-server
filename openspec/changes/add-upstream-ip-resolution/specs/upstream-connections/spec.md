## ADDED Requirements

### Requirement: Upstream IP Resolution with Fallback

The system SHALL provide the ability to resolve upstream proxy domain names to
IP addresses using configured DNS/DoH services and attempt sequential
connections to resolved IP addresses until a successful connection is
established.

#### Scenario: Upstream IP resolution enabled with successful connection

- **WHEN** `-upstream-resolve-ips` command-line parameter is enabled and an
  upstream proxy is configured with a domain name
- **AND** DNS/DoH resolution returns multiple IP addresses for the upstream
  domain
- **THEN** the system SHALL attempt to connect to each IP address in sequence
  until one succeeds
- **AND** successful connection details SHALL be logged with the connected IP
  address

#### Scenario: Upstream IP resolution with connection failures

- **WHEN** `-upstream-resolve-ips` command-line parameter is enabled and the
  first few resolved IP addresses fail to connect
- **THEN** the system SHALL continue attempting connection to remaining resolved
  IP addresses
- **AND** all connection attempt failures SHALL be logged with detailed error
  information
- **AND** the system SHALL return an error only after all resolved IP addresses
  have been exhausted

#### Scenario: DNS resolution failure for upstream

- **WHEN** `-upstream-resolve-ips` command-line parameter is enabled and DNS/DoH
  resolution fails to resolve the upstream domain
- **THEN** the system SHALL log the resolution failure and fall back to
  domain-based connection
- **AND** SHALL continue with existing upstream connection behavior

#### Scenario: Backward compatibility

- **WHEN** `-upstream-resolve-ips` command-line parameter is not enabled
  (default behavior)
- **THEN** the system SHALL use existing domain-based upstream connection logic
- **AND** SHALL not perform any IP resolution for upstream connections

## MODIFIED Requirements

### Requirement: Upstream Connection Configuration

The existing upstream proxy connection configuration SHALL be extended to
support optional IP-based connection mode while maintaining full backward
compatibility with domain-based connections.

The system SHALL:

- Accept new `-upstream-resolve-ips` boolean command-line parameter defaulting
  to false
- When enabled, resolve upstream domain names to IP addresses using the
  configured DNS/DoH infrastructure
- Attempt connections to resolved IP addresses in random order to distribute
  load
- Provide detailed logging of connection attempts including IP addresses,
  success/failure status, and timing information
- Maintain all existing upstream authentication and configuration parameters for
  each IP connection attempt
- Fall back to domain-based connection if IP resolution fails or is unavailable

#### Scenario: IP resolution disabled (backward compatibility)

- **WHEN** `-upstream-resolve-ips` is false or not specified
- **THEN** the system SHALL use existing domain-based upstream connection logic
- **AND** SHALL maintain all current functionality without any IP resolution

#### Scenario: IP resolution enabled with multiple upstream types

- **WHEN** `-upstream-resolve-ips` is true and multiple upstream proxy types are
  configured (HTTP, SOCKS5, WebSocket)
- **THEN** the system SHALL resolve domain names for all upstream types to IP
  addresses
- **AND** SHALL apply IP resolution to each upstream type independently with
  their respective authentication parameters
