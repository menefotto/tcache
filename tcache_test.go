package tcache

import (
	"testing"
	"time"
)

func TestCachePutGet(t *testing.T) {
	tc := New(1, 2)
	defer tc.Stop()

	tc.Put("key", "value")
	val, ok := tc.Get("key")
	if !ok {
		t.Fatal("Not ok fatal!")
	}

	if val.(string) != "value" {
		t.Fatal("Should be equal")
	}
}

func TestExpiration(t *testing.T) {
	tc := New(1, 2)
	defer tc.Stop()

	tc.Put("key", "value")

	time.Sleep(3 * time.Minute)

	_, ok := tc.Get("key")
	if ok {
		t.Fatal("Not not be ok, fatal!")
	}
}
