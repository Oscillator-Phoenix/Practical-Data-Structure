package lru

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	nSimpleKVs int = 6
)

func simpleNumKVs() []keyValue {
	kvs := []keyValue{}

	for i := 0; i < nSimpleKVs; i++ {
		s := strconv.Itoa(i)
		kvs = append(kvs, keyValue{s, s})
	}

	return kvs
}

func simpleAlphaKVs() []keyValue {
	kvs := []keyValue{}

	for i := 0; i < nSimpleKVs; i++ {
		s := string('a' + i)
		kvs = append(kvs, keyValue{s, s})
	}

	return kvs
}

func TestDoubleListPushFront(t *testing.T) {
	dl := newDoubleList()
	kvs := simpleNumKVs()
	for _, kv := range kvs {
		dl.pushFront(kv)
	}
	fmt.Println(dl)
}

func TestDoubleListPushBack(t *testing.T) {
	dl := newDoubleList()
	kvs := simpleNumKVs()
	for _, kv := range kvs {
		dl.pushBack(kv)
	}
	fmt.Println(dl)
}

func TestDoubleListInsertBefore(t *testing.T) {
	dl := newDoubleList()
	numKvs := simpleNumKVs()
	alphaKVs := simpleAlphaKVs()
	for i := 0; i < nSimpleKVs; i++ {
		e := dl.pushBack(numKvs[i])
		dl.insertBefore(alphaKVs[i], e)
	}
	fmt.Println(dl)
}

func TestDoubleListInsertAfter(t *testing.T) {
	dl := newDoubleList()
	numKvs := simpleNumKVs()
	alphaKVs := simpleAlphaKVs()
	for i := 0; i < nSimpleKVs; i++ {
		e := dl.pushBack(numKvs[i])
		dl.insertAfter(alphaKVs[i], e)
	}
	fmt.Println(dl)
}

func TestDoubleListRemove(t *testing.T) {
	dl := newDoubleList()
	numKvs := simpleNumKVs()
	alphaKVs := simpleAlphaKVs()
	toRemoves := [](*element){}

	for i := 0; i < nSimpleKVs; i++ {
		numEle := dl.pushBack(numKvs[i])
		alphaEle := dl.insertAfter(alphaKVs[i], numEle)
		toRemoves = append(toRemoves, alphaEle)
	}
	fmt.Println("Before removing: ", dl)

	for _, tr := range toRemoves {
		dl.remove(tr)
	}
	fmt.Println("After removing: ", dl)
}

func TestDoubleListMoveToFront(t *testing.T) {
	dl := newDoubleList()
	numKvs := simpleNumKVs()
	alphaKVs := simpleAlphaKVs()
	toFronts := [](*element){}

	for i := 0; i < nSimpleKVs; i++ {
		numEle := dl.pushBack(numKvs[i])
		alphaEle := dl.insertAfter(alphaKVs[i], numEle)
		toFronts = append(toFronts, alphaEle)
	}
	fmt.Println("Before moving to front: ", dl)

	for _, tr := range toFronts {
		dl.moveToFront(tr)
	}
	fmt.Println("After moving to front: ", dl)
}

func TestDoubleListMoveToBack(t *testing.T) {
	dl := newDoubleList()
	numKvs := simpleNumKVs()
	alphaKVs := simpleAlphaKVs()
	toBacks := [](*element){}

	for i := 0; i < nSimpleKVs; i++ {
		numEle := dl.pushBack(numKvs[i])
		alphaEle := dl.insertAfter(alphaKVs[i], numEle)
		toBacks = append(toBacks, alphaEle)
	}
	fmt.Println("Before moving to back: ", dl)

	for _, tr := range toBacks {
		dl.moveToBack(tr)
	}
	fmt.Println("After moving to back: ", dl)
}
