package binsearchtree

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var testLess lessFunc = func(key1, key2 int) bool { return key1 < key2 }

func TestBST1(t *testing.T) {
	// bst := newBinarySearchTree(testLess)
	// fmt.Println(bst)
	// bst.insert(3)
	// bst.insert(5)
	// bst.insert(2)
	// bst.insert(1)
	// bst.insert(4)
	// bst.insert(3)
	// fmt.Println(bst)
	// fmt.Println(bst.sortedVals())
}

type sortFunc func(xs []int, less lessFunc) []int

const (
	testTimes      = 100
	testScaleLimit = 10000
	testNumRange   = 100000
)

func genRandomInts(scaleLimit int) []int {
	rand.Seed(time.Now().Unix())
	scale := rand.Intn(scaleLimit)
	xs := make([]int, scale)
	for i := 0; i < scale; i++ {
		xs[i] = rand.Intn(testNumRange)
	}
	return xs
}

func sortOneCheck(mySort sortFunc) bool {
	var testLess lessFunc = func(key1, key2 int) bool { return key1 < key2 }
	xs := genRandomInts(testScaleLimit)

	predict := mySort(xs, testLess)

	answer := make([]int, len(xs))
	copy(answer, xs)
	sort.Slice(answer, func(i, j int) bool { return testLess(answer[i], answer[j]) })

	if len(predict) != len(xs) || len(answer) != len(xs) {
		fmt.Printf("len(predict) = %d, len(xs) = %d\n", len(predict), len(xs))
		return false
	}

	for i := 0; i < len(xs); i++ {
		if predict[i] != answer[i] {
			fmt.Printf("predict[%d] = %d, answer[%d] = %d\n", i, predict[i], i, answer[i])
			return false
		}
	}

	return true
}

func sortCheck(mySort sortFunc) bool {
	for i := 0; i < testTimes; i++ {
		if sortOneCheck(mySort) == false {
			return false
		}
	}
	return true
}

func TestBST2(t *testing.T) {
	sf := func(xs []int, less lessFunc) []int {
		bst := newBinarySearchTree(less)
		for _, x := range xs {
			bst.insert(x)
		}
		return bst.sortedVals()
	}
	if sortCheck(sf) != true {
		t.Fail()
	}
}
