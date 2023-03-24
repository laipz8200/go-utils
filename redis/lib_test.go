package cache

import (
	"context"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	cache := NewRedisCache("localhost:32768", "redispw", 0)

	// Test Set and Get
	testTable := []struct {
		key   string
		value string
	}{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}

	for _, tt := range testTable {
		err := cache.Set(context.Background(), tt.key, tt.value, time.Minute)
		if err != nil {
			t.Errorf("failed to set key %s: %v", tt.key, err)
		}

		val, err := cache.Get(context.Background(), tt.key)
		if err != nil {
			t.Errorf("failed to get key %s: %v", tt.key, err)
		}

		if val != tt.value {
			t.Errorf("expected value %s, got %s", tt.value, val)
		}
	}

	// Test Delete
	for _, tt := range testTable {
		err := cache.Delete(context.Background(), tt.key)
		if err != nil {
			t.Errorf("failed to delete key %s: %v", tt.key, err)
		}

		val, err := cache.Get(context.Background(), tt.key)
		if err != nil {
			t.Errorf("failed to get key %s: %v", tt.key, err)
		}

		if val != "" {
			t.Errorf("expected value '', got %s", val)
		}
	}
}

func TestRedisCache_ErrorCases(t *testing.T) {
	ctx := context.Background()
	cache := NewRedisCache("localhost:32768", "redispw", 0)

	// Test Get error case
	_, err := cache.Get(ctx, "non_existing_key")
	if err != nil {
		t.Errorf("Expected nil error, but got %v", err)
	}

	// Test Set error case
	err = cache.Set(ctx, "", "value", time.Minute)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test Delete error case
	err = cache.Delete(ctx, "")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
