package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(1 * time.Second)

	if cache == nil {
		t.Error("Expected non-nil cache, got nil")
	}
}

func TestDeleteStale(t *testing.T) {
	cache := NewCache(1 * time.Second)

	key1 := "staleKey1"
	key2 := "staleKey2"

	cache.Add(key1, []byte("staleValue1"))
	cache.Add(key2, []byte("staleValue2"))

	// Sleep to allow the entries to become stale
	time.Sleep(2 * time.Second)

	cache.deleteStale(1 * time.Second)

	_, existsKey1 := cache.Get(key1)
	_, existsKey2 := cache.Get(key2)

	if existsKey1 || existsKey2 {
		t.Error("Expected stale entries to be deleted, but they still exist")
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
