package lru

import "fmt"

type keyValue struct {
	key   string
	value string
}

func (kv keyValue) String() string {
	return fmt.Sprintf("{%s}", kv.key)
	// return fmt.Sprintf("{%s : %s}", kv.key, kv.value)
}

// doubleList

type lruCache struct {
	m map[string]*element
}
