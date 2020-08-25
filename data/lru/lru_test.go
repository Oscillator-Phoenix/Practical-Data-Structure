package lru

import (
	"fmt"
	"testing"
)

func genCase1() []keyValue {
	return []keyValue{
		{"key1", "1"},
		{"key2", "2"},
	}
}

func genCase2() []keyValue {
	return []keyValue{
		{"key1", "1"},
		{"key2", "2"},
		{"k3", "3"},
		{"k3", "4"},
	}
}

func TestCacheGet(t *testing.T) {
	var tests = genCase1()
	lru := newLRUCache(0) // unlimited

	for _, tt := range tests {
		lru.put(tt.key, tt.value)
	}

	for i, tt := range tests {
		if v, err := lru.get(tt.key); err != nil || v != tt.value {
			t.Fatalf("%d test failed. got %v ,wanted %v", i, v, tt.value)
		}
	}
	fmt.Println(lru.dl)
}

func TestRemoveOldest(t *testing.T) {
	var tests = genCase2()
	limit := tests[0].nBytes() + tests[1].nBytes()
	lru := newLRUCache(limit)

	for _, tt := range tests {
		lru.put(tt.key, tt.value)
	}
	fmt.Println(lru.dl)

	if _, err := lru.get(tests[0].key); err == nil {
		t.Fatalf("Removeoldest key1 failed")
	}
	if lru.size() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}
