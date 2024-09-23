package cache

import (
	"errors"
	"testing"
	"time"
)

func TestGetMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := 1
	fakeKey := Key("name")

	c.Set(testKey, testValue, 2)

	actualValue, err := c.Get(testKey)
	if err != nil {
		t.Errorf("Get(%v) returned error: %v, want: %v\n", testKey, err, testValue)
	}
	if actualValue != testValue {
		t.Errorf("Get(%v) returned: %v, want: %v\n", testKey, actualValue, testValue)
	}

	_, err = c.Get(fakeKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) returned: %v, want: %v\n", fakeKey, err, errKeyNotFound)
	}
}

func TestSetMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := Value(1)

	c.Set(testKey, testValue, 2)

	_, err := c.Get(testKey)
	if err != nil {
		t.Errorf("Get(%v) returned error: %v, want: nil\n", testKey, err)
	}
}

func TestDeleteMethod(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := Value(1)
	fakeKey := Key("name")

	c.Set(testKey, testValue, 2)

	err := c.Delete(testKey)
	if err != nil {
		t.Errorf("Delete(%v) returned: %v, want: nill\n", testKey, err)
	}

	_, err = c.Get(testKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) after Delete(%v) returned: %v, wait: %v\n", testKey, testKey, err, errKeyNotFound)
	}

	err = c.Delete(fakeKey)
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Delete(%v) returned: %v, wait: %v\n", fakeKey, err, errKeyNotFound)
	}
}

func TestLifetime(t *testing.T) {
	c := New()

	testKey := Key("userId")
	testValue := 77

	c.Set(testKey, testValue, 2)

	actualValue, _ := c.Get("userId")
	if actualValue != testValue {
		t.Errorf("Get(%v) returned: %v, wait: %v\n", testKey, actualValue, testValue)
	}

	time.Sleep(time.Second * 3)

	_, err := c.Get("userId")
	if !errors.Is(err, errKeyNotFound) {
		t.Errorf("Get(%v) returned: %v, wait: %v\n", testKey, err, errKeyNotFound)
	}
}
