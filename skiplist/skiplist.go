package skiplist

import (
	"errors"
	"math/rand"
)

// skipList is an ordered key-value map which was proposed by the paper below:
// https://www.epaperpress.com/sortsearch/download/skiplist.pdf

const (
	defaultMaxLevel    int     = 32
	defaultProbability float32 = 0.25
)

type lessFunc func(key1, key2 int) bool

var defaultLessFunc lessFunc = func(key1, key2 int) bool { return key1 < key2 }

type keyValue struct {
	key   int
	value int
}

type skipListNode struct {
	keyValue
	forwards [](*skipListNode) // length of `forwards` is the level of this node
}

type skipList struct {
	head     *skipListNode
	_size    int
	maxLevel int
	p        float32
	less     lessFunc
}

func newSkipListNode(key, value int, level int) *skipListNode {
	var node skipListNode
	node.key = key
	node.value = value
	node.forwards = make([](*skipListNode), level)
	return &node
}

func newSkipList() skipList {
	var sl skipList
	sl._size = 0
	sl.p = defaultProbability
	sl.maxLevel = defaultMaxLevel
	sl.less = defaultLessFunc
	sl.head = newSkipListNode(0, 0, sl.maxLevel) // initialize head-node with maxLevel
	return sl
}

// randomLevel returns a random level according to `p` and `maxLevel`
func (sl *skipList) randomLevel() int {
	level := 1
	for rand.Float32() < sl.p && level < sl.maxLevel {
		level++
	}
	return level
}

func (sl *skipList) size() int {
	return sl._size
}

func (sl *skipList) empty() bool {
	return sl.size() == 0
}

func (sl *skipList) search(key int) (value int, ok bool) {
	x := sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for x.forwards[i] != nil && sl.less(x.forwards[i].key, key) {
			x = x.forwards[i] // skip
		}
	}
	x = x.forwards[0]
	if x != nil && x.key == key {
		return x.value, true
	}
	return 0, false
}

func (sl *skipList) insert(key, value int) error {
	// fmt.Println("head", sl.head)
	// fmt.Printf("to insert: (%d, %d)\n", key, value)

	update := make([]*skipListNode, sl.maxLevel)
	x := sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for x.forwards[i] != nil && sl.less(x.forwards[i].key, key) {
			x = x.forwards[i] // skip
		}
		update[i] = x
	}
	x = x.forwards[0]

	// fmt.Println("search done")

	// replace existing old value with the new value, then return
	if x != nil && x.key == key {
		x.value = value
		return nil // insert succeeded
	}

	newNodeLevel := sl.randomLevel() // function `randomLevel` make sure `newNodeLevel < sl.maxLevel`
	// fmt.Println("newNodeLevel", newNodeLevel)

	newNode := newSkipListNode(key, value, newNodeLevel)
	for i := 0; i < newNodeLevel; i++ {
		newNode.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = newNode
	}
	sl._size++

	// fmt.Println("inserted (", key, ", ", value, ")")
	return nil // insert succeeded
}

func (sl *skipList) delete(key int) error {
	if sl.empty() {
		return errors.New("can not detele on an empty skip list")
	}

	update := make([]*skipListNode, sl.maxLevel)
	x := sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for x.forwards[i] != nil && sl.less(x.forwards[i].key, key) {
			x = x.forwards[i] // skip
		}
		update[i] = x
	}

	x = x.forwards[0]

	if x != nil && x.key == key {
		for i := 0; i < sl.maxLevel; i++ {
			if update[i].forwards[i] != x {
				return nil // level of x done
			}
			update[i].forwards[i] = x.forwards[i]
		}
		sl._size--
	}

	return nil
}

// traverse traverses the skipList in the order defined by lessFunc
func (sl *skipList) traverse(operate func(key, value int)) {
	// itereate on level-0 which is a single linked list
	x := sl.head.forwards[0]
	for x != nil {
		operate(x.key, x.value)
		x = x.forwards[0]
	}
}
