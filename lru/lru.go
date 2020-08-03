package lru

import (
	"errors"
	"fmt"
)

type keyValue struct {
	key   string
	value string
}

func (kv keyValue) String() string {
	// return fmt.Sprintf("{%s : %s}", kv.key, kv.value)
	return fmt.Sprintf("{%s}", kv.key)
}

func (kv keyValue) nBytes() int {
	return len(kv.key) + len(kv.value)
}

// lruCache is a key-value cache using Least-Recently-Used replacement strategy
type lruCache struct {
	_bytesLimit int                   // limit of bytes the cache could cost, unlimited when `_bytesLimit <= 0`
	_nBytes     int                   // number of bytes the cache has cost
	_size       int                   // number of reconds in cache
	m           map[string](*element) // hash-map to implement LRU
	dl          *doublelist           // double-linked-list to implement LRU
}

func newLRUCache(bytesLimit int) *lruCache {
	var c lruCache

	c._size = 0
	c._nBytes = 0
	c._bytesLimit = bytesLimit
	c.m = make(map[string](*element))
	c.dl = newDoubleList()

	return &c
}

func (c *lruCache) size() int {
	return c._size
}

func (c *lruCache) empty() bool {
	return c.size() == 0
}

func (c *lruCache) nBytes() int {
	return c._nBytes
}

func (c *lruCache) bytesLimit() int {
	return c._bytesLimit
}

func (c *lruCache) setBytesLimit(limit int) {
	c._bytesLimit = limit
}

func (c *lruCache) get(key string) (value string, err error) {
	if ele, isPresent := c.m[key]; isPresent {
		c.dl.moveToFront(ele) // move to head
		kv := ele.value       // kv is just a pointer
		return kv.value, nil
	}
	return value, errors.New("absent key")
}

func (c *lruCache) removeOldest() {
	// the bakc of list is oldest element in cache
	if ele := c.dl.back(); ele != nil {
		c.dl.remove(ele)
		kv := ele.value // kv is just a pointer
		delete(c.m, kv.key)
		c._nBytes -= (len(kv.key) + len(kv.value))
		c._size--
	}
}

func (c *lruCache) put(key, value string) (err error) {
	if ele, isPresent := c.m[key]; isPresent {
		// if kv is present, just update the value of kv
		c.dl.moveToFront(ele)
		kv := ele.value // kv is just a pointer
		c._nBytes += (len(value) - len(kv.value))
		kv.value = value
	} else {
		ele := c.dl.pushFront(&keyValue{key, value})
		c.m[key] = ele
		c._nBytes += (len(key) + len(value))
		c._size++
	}
	if c._bytesLimit > 0 {
		for c._nBytes > c._bytesLimit {
			c.removeOldest()
		}
	}
	return nil
}
