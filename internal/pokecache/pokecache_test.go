package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected item with key %s to be in cache", key)
	}
	if string(retrievedValue) != string(value) {
		t.Fatalf("expected value %s, got %s", value, retrievedValue)
	}
}

func TestCacheItemExpiration(t *testing.T) {
	interval := 2 * time.Second
	cache := NewCache(interval)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	time.Sleep(3 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Fatalf("expected item with key %s to be expired from cache", key)
	}
}

func TestCacheReapLoop(t *testing.T) {
	interval := 2 * time.Second
	cache := NewCache(interval)

	key1 := "testKey1"
	value1 := []byte("testValue1")
	key2 := "testKey2"
	value2 := []byte("testValue2")

	cache.Add(key1, value1)
	cache.Add(key2, value2)

	time.Sleep(1 * time.Second)
	if _, ok := cache.Get(key1); !ok {
		t.Fatalf("expected item with key %s to be in cache", key1)
	}
	if _, ok := cache.Get(key2); !ok {
		t.Fatalf("expected item with key %s to be in cache", key2)
	}

	time.Sleep(2 * time.Second)
	if _, ok := cache.Get(key1); ok {
		t.Fatalf("expected item with key %s to be expired from cache", key1)
	}
	if _, ok := cache.Get(key2); ok {
		t.Fatalf("expected item with key %s to be expired from cache", key2)
	}
}
