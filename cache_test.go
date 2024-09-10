package cache

import (
	"errors"
	"testing"
)

func TestGetMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := Value(1)
	fakeKey := Key("name")

	c.Set(testKey, testValue)

	actualValue, err := c.Get(testKey)
	if err != nil {
		t.Errorf("Get(%v) returned error: %v, want: %v", testKey, err, testValue)
	}
	if actualValue != testValue {
		t.Errorf("Get(%v) returned: %v, want: %v", testKey, actualValue, testValue)
	}

	_, err = c.Get(fakeKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) returned: %v, want: %v", fakeKey, err, errKeyNotFound)
	}
}

func TestSetMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := Value(1)

	c.Set(testKey, testValue)

	_, err := c.Get(testKey)
	if err != nil {
		t.Errorf("Get(%v) returned error: %v, want: nil", testKey, err)
	}
}

func TestDeleteMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := Value(1)
	fakeKey := Key("name")

	c.Set(testKey, testValue)

	err := c.Delete(testKey)
	if err != nil {
		t.Errorf("Delete(%v) returned: %v, want: nill", testKey, err)
	}

	_, err = c.Get(testKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) after Delete(%v) returned: %v, wait: %v", testKey, testKey, err, errKeyNotFound)
	}

	err = c.Delete(fakeKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Delete(%v) returned: %v, wait: %v", fakeKey, err, errKeyNotFound)
	}
}
