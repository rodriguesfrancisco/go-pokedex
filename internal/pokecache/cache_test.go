package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)
	if cache == nil {
		t.Error("NewCache returned nil")
	}
	if cache.entries == nil {
		t.Error("cache entries map is nil")
	}
	if cache.interval != time.Millisecond*100 {
		t.Errorf("expected interval %v, got %v", time.Millisecond*100, cache.interval)
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(time.Minute)

	t.Run("add and retrieve value", func(t *testing.T) {
		key := "test-key"
		val := []byte("test-value")

		cache.Add(key, val)
		retrieved, ok := cache.Get(key)

		if !ok {
			t.Error("expected key to exist in cache")
		}

		if string(retrieved) != string(val) {
			t.Errorf("expected value %s, got %s", string(val), string(retrieved))
		}
	})

	t.Run("get non-existent key", func(t *testing.T) {
		_, ok := cache.Get("non-existent")
		if ok {
			t.Error("expected key to not exist in cache")
		}
	})

	t.Run("overwrite existing key", func(t *testing.T) {
		key := "overwrite-key"
		val1 := []byte("first-value")
		val2 := []byte("second-value")

		cache.Add(key, val1)
		cache.Add(key, val2)

		retrieved, ok := cache.Get(key)
		if !ok {
			t.Error("expected key to exist in cache")
		}

		if string(retrieved) != string(val2) {
			t.Errorf("expected value %s, got %s", string(val2), string(retrieved))
		}
	})
}

func TestCacheReap(t *testing.T) {
	t.Run("entries are reaped after interval", func(t *testing.T) {
		interval := time.Millisecond * 100
		cache := NewCache(interval)

		key := "reap-test"
		val := []byte("will-be-reaped")

		cache.Add(key, val)

		// Verify entry exists
		_, ok := cache.Get(key)
		if !ok {
			t.Error("expected key to exist initially")
		}

		// Wait for slightly more than the interval plus reap loop interval
		time.Sleep(interval + interval + time.Millisecond*50)

		// Verify entry was reaped
		_, ok = cache.Get(key)
		if ok {
			t.Error("expected key to be reaped after interval")
		}
	})

	t.Run("recent entries are not reaped", func(t *testing.T) {
		interval := time.Second
		cache := NewCache(interval)

		key := "keep-test"
		val := []byte("will-be-kept")

		cache.Add(key, val)

		// Wait for less than the interval
		time.Sleep(time.Millisecond * 100)

		// Verify entry still exists
		retrieved, ok := cache.Get(key)
		if !ok {
			t.Error("expected key to still exist")
		}
		if string(retrieved) != string(val) {
			t.Errorf("expected value %s, got %s", string(val), string(retrieved))
		}
	})
}

func TestCacheMultipleEntries(t *testing.T) {
	cache := NewCache(time.Minute)

	entries := map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
	}

	// Add all entries
	for key, val := range entries {
		cache.Add(key, val)
	}

	// Verify all entries exist
	for key, expectedVal := range entries {
		retrieved, ok := cache.Get(key)
		if !ok {
			t.Errorf("expected key %s to exist", key)
		}
		if string(retrieved) != string(expectedVal) {
			t.Errorf("for key %s: expected value %s, got %s", key, string(expectedVal), string(retrieved))
		}
	}
}

func TestCacheConcurrency(t *testing.T) {
	cache := NewCache(time.Minute)

	// Test concurrent writes and reads
	done := make(chan bool)

	// Writer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Add("concurrent-key", []byte("concurrent-value"))
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Get("concurrent-key")
		}
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done
}

func TestCacheEmptyValue(t *testing.T) {
	cache := NewCache(time.Minute)

	key := "empty-key"
	val := []byte{}

	cache.Add(key, val)
	retrieved, ok := cache.Get(key)

	if !ok {
		t.Error("expected key to exist in cache")
	}

	if len(retrieved) != 0 {
		t.Errorf("expected empty value, got %v", retrieved)
	}
}

func TestCacheNilValue(t *testing.T) {
	cache := NewCache(time.Minute)

	key := "nil-key"
	var val []byte = nil

	cache.Add(key, val)
	retrieved, ok := cache.Get(key)

	if !ok {
		t.Error("expected key to exist in cache")
	}

	if retrieved != nil {
		t.Errorf("expected nil value, got %v", retrieved)
	}
}
