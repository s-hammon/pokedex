package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	tests := []struct {
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

	for _, tt := range tests {
		t.Run(fmt.Sprintf("key=%s", tt.key), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tt.key, tt.val)
			val, ok := cache.Get(tt.key)
			if !ok {
				t.Errorf("expected to find key %q", tt.key)
			}
			if string(val) != string(tt.val) {
				t.Errorf("expected %q, got %q", tt.val, val)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)

	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected key to be removed")
	}
}
