package lru

import (
	"testing"
)

func TestCache_Get(t *testing.T) {
	lru := New(int64(0))
	lru.Add("key1", []byte("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v) != "1234" {
		t.Fatal("cache hit key1=1234 fail")
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatal("cache miss key2 fail")
	}
}

func TestCache_Add(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	caps := len(k1 + k2 + v1 + v2)
	lru := New(int64(caps))
	lru.Add(k1, []byte(v1))
	lru.Add(k2, []byte(v2))
	lru.Add(k3, []byte(v3))

	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
