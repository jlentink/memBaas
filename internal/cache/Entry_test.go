package cache

import (
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	key := "testKey"
	value := "testValue"
	ttl := int64(60)

	entry := NewEntry(key, value, ttl)

	if entry.Key != key {
		t.Errorf("expected Key to be %s, got %s", key, entry.Key)
	}
	if entry.Data != value {
		t.Errorf("expected Data to be %s, got %s", value, entry.Data)
	}
	if entry.TTL != ttl {
		t.Errorf("expected TTL to be %d, got %d", ttl, entry.TTL)
	}
	if time.Since(entry.Created) > time.Second {
		t.Errorf("expected Created to be recent, got %s", entry.Created)
	}
}
