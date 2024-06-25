package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Fatalf("expected non-nil map, got nil")
	}
}

func TestAdd(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	val := []byte("testValue")
	cache.Add(key, val)

	entry, exists := cache.cache[key]
	if !exists {
		t.Fatalf("expected entry to exist for key %s", key)
	}
	if string(entry.val) != string(val) {
		t.Fatalf("expected value %s, got %s", val, entry.val)
	}
}

func TestGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	val := []byte("testValue")
	cache.Add(key, val)

	retrievedVal, exists := cache.Get(key)
	if !exists {
		t.Fatalf("expected entry to exist for key %s", key)
	}
	if string(retrievedVal) != string(val) {
		t.Fatalf("expected value %s, got %s", val, retrievedVal)
	}
}

func TestGetNonExistent(t *testing.T) {
	cache := NewCache()
	_, exists := cache.Get("nonExistentKey")
	if exists {
		t.Fatalf("expected entry to not exist")
	}
}

func TestCacheEntryTimestamp(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	val := []byte("testValue")
	startTime := time.Now().UTC()
	cache.Add(key, val)

	entry, exists := cache.cache[key]
	if !exists {
		t.Fatalf("expected entry to exist for key %s", key)
	}

	if entry.createdAt.Before(startTime) || entry.createdAt.After(time.Now().UTC()) {
		t.Fatalf("expected timestamp to be between %v and %v, got %v", startTime, time.Now().UTC(), entry.createdAt)
	}
}
