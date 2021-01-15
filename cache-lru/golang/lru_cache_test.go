package main

import "testing"

func TestGet(t *testing.T) {
	cache := NewLRUCache(3)
	if err := cache.Get(1); err == nil {
		t.Errorf("Get() on empty cache should return error")
	}

	cache.Add(1)
	if err := cache.Get(1); err != nil {
		t.Errorf("Get() on existing element should return err = nil")
	}
}

func TestAddPolicy(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Add(1)
	cache.Add(2)
	cache.Add(3)
	cache.Add(4)
	// expected state: [2, 3, 4]

	if err := cache.Get(1); err == nil {
		t.Errorf("Get() on expired value should return error")
	}
	err2 := cache.Get(2)
	err3 := cache.Get(3)
	err4 := cache.Get(4)
	if err2 != nil || err3 != nil || err4 != nil {
		t.Errorf("Get() on present values should not return error")
	}

	cache.Add(5)
	// expected state: [3, 4, 5]
	if err := cache.Get(2); err == nil {
		t.Errorf("Get() on expired value should return error")
	}
	err3 = cache.Get(3)
	err4 = cache.Get(4)
	err5 := cache.Get(5)
	if err3 != nil || err4 != nil || err5 == nil {
		t.Errorf("Get() on expired value should return error")
	}

	cache.Add(1)
	// expected state: [4, 5, 1]
	if err := cache.Get(2); err == nil {
		t.Errorf("Get() on present values should not return error")
	}
	err4 = cache.Get(4)
	err5 = cache.Get(5)
	err1 := cache.Get(5)
	if err4 != nil || err5 != nil || err1 == nil {
		t.Errorf("Get() on expired value should return error")
	}

}
