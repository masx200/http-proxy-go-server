## ADDED Requirements

### Requirement: DNS Resolution Caching
The system SHALL cache DNS query results to improve performance and reduce external DNS service load.

#### Scenario: Cache hit for domain resolution
- **WHEN** a domain is requested that exists in the cache
- **THEN** the cached IP address SHALL be returned without performing external DNS resolution
- **AND** the cache entry SHALL be valid (not expired)

#### Scenario: Cache miss for domain resolution
- **WHEN** a domain is requested that does not exist in the cache or cache entry is expired
- **THEN** external DNS resolution SHALL be performed
- **AND** the result SHALL be stored in the cache with appropriate TTL
- **AND** the resolved IP address SHALL be returned

### Requirement: Cache Persistence
The DNS cache SHALL persist to local file system and survive application restarts.

#### Scenario: Application startup with existing cache
- **WHEN** the application starts and a cache file exists
- **THEN** the system SHALL load valid cache entries from the file
- **AND** expired entries SHALL be discarded during load
- **AND** successful cache load SHALL not block application startup

#### Scenario: Application shutdown
- **WHEN** the application shuts down gracefully
- **THEN** all current cache entries SHALL be saved to the persistent file
- **AND** the save operation SHALL be atomic (temp file + rename)

### Requirement: Cache Configuration via Command Line
The system SHALL support cache configuration through command-line parameters.

#### Scenario: Custom cache file path
- **WHEN** `-cache-file` parameter is provided
- **THEN** the system SHALL use the specified file path for cache storage
- **AND** SHALL create necessary directories if they don't exist

#### Scenario: Custom TTL configuration
- **WHEN** `-cache-ttl` parameter is provided
- **THEN** the system SHALL use the specified TTL duration for cache entries
- **AND** SHALL accept duration strings (e.g., "5m", "10m", "1h")

#### Scenario: Custom save interval
- **WHEN** `-cache-save-interval` parameter is provided
- **THEN** the system SHALL save cache to file at the specified interval
- **AND** SHALL continue to serve requests during save operations

### Requirement: Cache Key Management
The cache SHALL use composite keys based on DNS record type and domain name.

#### Scenario: Different DNS record types
- **WHEN** resolving A records and AAAA records for the same domain
- **THEN** each record type SHALL be cached separately
- **AND** cache keys SHALL follow the format "TYPE:domain" (e.g., "A:example.com")

#### Scenario: Domain name normalization
- **WHEN** caching entries for domains with different formats (example.com, example.com.)
- **THEN** domain names SHALL be normalized before cache key generation
- **AND** trailing dots SHALL be removed
- **AND** case SHALL be standardized to lowercase

### Requirement: Cache Integration with Resolvers
DNS caching SHALL be transparently integrated into all existing resolver implementations.

#### Scenario: Hosts resolver with caching
- **WHEN** using HostsResolver
- **THEN** hosts file resolution results SHALL be cached
- **AND** subsequent resolutions SHALL use cached results when available

#### Scenario: DOH resolver with caching
- **WHEN** using DOHResolver or DOH3Resolver
- **THEN** DoH resolution results SHALL be cached
- **AND** cache SHALL reduce DoH server load
- **AND** SHALL improve response times for cached entries

#### Scenario: Combined resolver with caching
- **WHEN** using HostsAndDohResolver
- **THEN** results from both hosts and DoH resolution SHALL be cached
- **AND** cache SHALL respect the resolver priority order (hosts first, then DoH)