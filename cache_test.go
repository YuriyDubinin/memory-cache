package cache

import (
	"errors"
	"testing"
)

func TestCache(t *testing.T) {
	c := New()

	// Test Set and Get
	key1 := Key("key1")
	value1 := "value1"

	c.Set(key1, value1)

	got, err := c.Get(key1)
	if err != nil {
		t.Fatalf("Get(%v) returned error: %v; want nil", key1, err)
	}
	if got != value1 {
		t.Errorf("Get(%v) = %v; want %v", key1, got, value1)
	}

	// Test Get with non-existent key
	key2 := Key("key2")
	_, err = c.Get(key2)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) returned error: %v; want %v", key2, err, errKeyNotFound)
	}

	// Test Delete
	err = c.Delete(key1)
	if err != nil {
		t.Fatalf("Delete(%v) returned error: %v; want nil", key1, err)
	}

	_, err = c.Get(key1)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) after Delete returned error: %v; want %v", key1, err, errKeyNotFound)
	}

	// Test Delete with non-existent key
	err = c.Delete(key2)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Delete(%v) returned error: %v; want %v", key2, err, errKeyNotFound)
	}
}
