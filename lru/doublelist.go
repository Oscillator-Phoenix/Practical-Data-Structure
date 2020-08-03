package lru

import (
	"fmt"
	"strings"
)

type element struct {
	value keyValue
	_prev *element
	_next *element
}

func (node *element) next() *element {
	return node._next
}

func (node *element) prev() *element {
	return node._prev
}

type doublelist struct {
	head  *element
	tail  *element
	_size int
}

func newDoubleList() *doublelist {
	var head element
	var tail element
	var list doublelist

	head._next = &tail
	tail._prev = &head
	list.head = &head
	list.tail = &tail
	list._size = 0

	return &list
}

func (dl *doublelist) size() int {
	return dl._size
}

func (dl *doublelist) empty() bool {
	return dl.size() == 0
}

func (dl *doublelist) front() *element {
	return dl.head._next
}

func (dl *doublelist) back() *element {
	return dl.tail._prev
}

// insertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified. The mark must not be nil.
func (dl *doublelist) insertAfter(v keyValue, mark *element) *element {
	x := dl.head

	for x != nil && x != mark {
		x = x._next
	}

	if x == nil {
		return nil
	}

	newElemet := &element{}
	newElemet.value = v
	newElemet._prev = mark
	newElemet._next = mark._next

	mark._next._prev = newElemet
	mark._next = newElemet

	return newElemet
}

func (dl *doublelist) insertBefore(v keyValue, mark *element) *element {
	x := dl.tail

	for x != nil && x != mark {
		x = x._prev
	}

	if x == nil {
		return nil
	}

	newElemet := &element{}
	newElemet.value = v
	newElemet._next = mark
	newElemet._prev = mark._prev

	mark._prev._next = newElemet
	mark._prev = newElemet

	return newElemet
}

func (dl *doublelist) pushBack(v keyValue) *element {
	return dl.insertBefore(v, dl.tail)
}

func (dl *doublelist) pushFront(v keyValue) *element {
	return dl.insertAfter(v, dl.head)
}

func (dl *doublelist) remove(e *element) {
	x := dl.head

	for x != nil && x != e {
		x = x._next
	}

	if x == e {
		// remove x
		x._prev._next = x._next
		x._next._prev = x._prev
	}
}

func (dl *doublelist) moveToBack(e *element) {
	x := dl.head

	for x != nil && x != e {
		x = x._next
	}

	if x == e {
		// remove x
		x._prev._next = x._next
		x._next._prev = x._prev
		// x to back
		x._next = dl.tail
		x._prev = dl.tail._prev
		dl.tail._prev._next = x
		dl.tail._prev = x
	}
}

func (dl *doublelist) moveToFront(e *element) {
	x := dl.head

	for x != nil && x != e {
		x = x._next
	}

	if x == e {
		// remove x
		x._prev._next = x._next
		x._next._prev = x._prev
		// x to back
		x._prev = dl.head
		x._next = dl.head._next
		dl.head._next._prev = x
		dl.head._next = x
	}
}

func (dl doublelist) String() string {
	var b strings.Builder
	b.WriteString("head <--> ")
	x := dl.head._next
	for x != dl.tail {
		fmt.Fprintf(&b, "%v <--> ", x.value)
		x = x._next
	}
	b.WriteString("tail")
	return b.String()
}
