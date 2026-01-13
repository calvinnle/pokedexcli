package pokecache

import (
	"testing"
	"time"
)

func TestCreatedCache(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)
	cache.Set("hello", []byte("world"))
	cache.Set("Lionel", []byte("Salamanca"))
	cache.Set("Gus", []byte("Fring"))

	cases := []struct {
		inputKey    string
		expectedOk  bool
		expectedVal string
	} {
		// correct value and ok
		{
			inputKey:    "hello",
			expectedOk:  true,
			expectedVal: "world",
		},

		{
			inputKey:    "Lionel",
			expectedOk:  true,
			expectedVal: "Salamanca",
		},
		{
			inputKey:    "Gus",
			expectedOk:  true,
			expectedVal: "Fring",
		},

		// cache miss
		{
			inputKey:    "biz",
			expectedOk:  false,
			expectedVal: "",
		},
	}

	for _, c := range cases {
		actualVal, actualOk := cache.Get(c.inputKey)
		if actualOk != c.expectedOk {
			t.Errorf("actual Ok: %v\n", actualOk)
			t.Errorf("expected Ok: %v\n", c.expectedOk)
		}

		if string(actualVal) != c.expectedVal {
			t.Errorf("actual val: %v\n", string(actualVal))
			t.Errorf("expected val: %v\n", c.expectedVal)
		}
	}
}

func TestReap(t *testing.T) {
    interval := time.Millisecond * 10
    cache := NewCache(interval)

    key1 := "key1"
    val1 := []byte("val1")

    cache.Set(key1, val1)

    time.Sleep(interval + (1 + time.Millisecond))

    cache.Set("key2", []byte("val2"))
    
    // Check key1 was deleted, key2 still exists
    if _, ok := cache.Get("key1"); ok {
        t.Error("Expected key1 to be reaped")
    }
    
    if _, ok := cache.Get("key2"); !ok {
        t.Error("Expected key2 to still exist")
    }
}


func TestReapFail(t *testing.T) {
    interval := time.Millisecond * 9
    cache := NewCache(interval)

    key1 := "key1"
    val1 := []byte("val1")

    cache.Set(key1, val1)

    time.Sleep(8 * time.Millisecond)

    cache.Set("key2", []byte("val2"))
    
    if _, ok := cache.Get("key1"); !ok {
        t.Error("key1 shouldn't have been reaped")
    }
}
