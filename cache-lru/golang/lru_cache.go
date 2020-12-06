package main

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	queue    *list.List
	mapping  map[int]*list.Element
}

type cacheEntry struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    list.New(),
		mapping:  make(map[int]*list.Element, capacity),
	}
}

func (c *LRUCache) Put(key, value int) {
	new := &cacheEntry{Key: key, Value: value}
	if c.queue.Len() == c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		asEntry := last.Value.(*cacheEntry)
		delete(c.mapping, asEntry.Key)
	}
	c.queue.PushFront(new)
	c.mapping[key] = c.queue.Front()
}

func (c *LRUCache) Get(key int) int {
	e, ok := c.mapping[key]
	if !ok {
		return -1
	}
	c.queue.Remove(e)
	c.queue.PushFront(e.Value)
	return e.Value.(*cacheEntry).Value
}
