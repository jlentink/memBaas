package cache

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		ttl       int64
		overwrite bool
		expected  Status
	}{
		{"key1", "value1", 60, false, Ok},
		{"key1", "value2", 60, false, NotUnique},
		{"key1", "value2", 60, true, OverwriteDenied},
		{"", "value", 60, false, InvalidKey},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			status := Set(tt.key, tt.value, tt.ttl, tt.overwrite)
			if status != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, status)
			}
		})
	}
}

func TestGet(t *testing.T) {
	Set("key1", "value1", 60, false)
	tests := []struct {
		key      string
		expected string
		status   Status
	}{
		{"key1", "value1", Ok},
		{"key2", "", NotFound},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			value, status := Get(tt.key)
			if value != tt.expected || status != tt.status {
				t.Errorf("expected %v and %v, got %v and %v", tt.expected, tt.status, value, status)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	Set("key1", "value1", 60, false)
	tests := []struct {
		key    string
		status Status
	}{
		{"key1", Ok},
		{"key2", NotFound},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			status := Delete(tt.key)
			if status != tt.status {
				t.Errorf("expected %v, got %v", tt.status, status)
			}
		})
	}
}

func TestCleanup(t *testing.T) {
	Set("key1", "value1", 1, false)
	time.Sleep(2 * time.Second)
	Cleanup()
	_, status := Get("key1")
	if status != NotFound {
		t.Errorf("expected %v, got %v", NotFound, status)
	}
}
