package dnscache

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// Simple debug test to understand go-cache behavior
func TestDebugCacheBehavior(t *testing.T) {
	tempDir := t.TempDir()
	cacheFile := filepath.Join(tempDir, "debug_cache.json")

	cache, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add a simple item
	testIP := net.ParseIP("192.168.1.1")
	cache.SetIPs("A", "example.com", []net.IP{testIP}, 5*time.Minute)

	// Debug: Check what's in the cache
	items := cache.cache.Items()
	fmt.Printf("Cache items count: %d\n", len(items))
	for k, v := range items {
		fmt.Printf("Key: %s, Expiration: %d, Object: %v\n", k, v.Expiration, v.Object)
	}

	// Save to file
	err = cache.Save()
	if err != nil {
		t.Fatalf("Failed to save cache: %v", err)
	}

	// Check if file exists and its content
	if data, err := os.ReadFile(cacheFile); err != nil {
		t.Fatalf("Failed to read cache file: %v", err)
	} else {
		fmt.Printf("Cache file content: %s\n", string(data))
	}

	cache.Close()

	// Create new cache and load
	cache2, err := New(cacheFile)
	if err != nil {
		t.Fatalf("Failed to create second cache: %v", err)
	}

	// Check what's loaded
	items2 := cache2.cache.Items()
	fmt.Printf("Loaded cache items count: %d\n", len(items2))
	for k, v := range items2 {
		fmt.Printf("Loaded Key: %s, Expiration: %d, Object: %v\n", k, v.Expiration, v.Object)
	}

	// Try to get the item directly from cache using the expected key format
	key := "A:example.com"
	if value, found := cache2.cache.Get(key); found {
		fmt.Printf("Raw value type: %T, value: %v\n", value, value)
		if ips, ok := value.([]net.IP); ok {
			fmt.Printf("Value is []net.IP: %v\n", ips)
		} else if ipsStr, ok := value.([]string); ok {
			fmt.Printf("Value is []string: %v\n", ipsStr)
		}
	}

	// Try to get the item using GetIPs
	loadedIPs, found := cache2.GetIPs("A", "example.com")
	fmt.Printf("Found IPs: %v, Found: %v\n", loadedIPs, found)

	cache2.Close()
}