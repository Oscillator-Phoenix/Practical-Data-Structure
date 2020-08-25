package binsearch

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	testTimes      = 100
	testScaleLimit = 10000
	testNumRange   = 100000
)

func genRandomIntsNoEmpty(scaleLimit int) []int {
	rand.Seed(time.Now().Unix())
	scale := rand.Intn(scaleLimit) + 1
	xs := make([]int, scale)
	for i := 0; i < scale; i++ {
		xs[i] = rand.Intn(testNumRange)
	}
	return xs
}

func TestBinarySearchRecursion1(t *testing.T) {
	if _, isFind := binarySearchRecursion([]int{}, 233); isFind != false {
		t.Fail()
	}
}

func TestBinarySearchRecursion2(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		xs := genRandomIntsNoEmpty(testScaleLimit)
		sort.Ints(xs)
		targetIdx := rand.Intn(len(xs))
		target := xs[targetIdx]
		idx, isFind := binarySearchRecursion(xs, target)
		if isFind != true || idx != targetIdx {
			fmt.Printf("target: xs[%d]=%d  found: xs[%d]=%d\n", targetIdx, target, idx, xs[idx])
			t.FailNow()
		}
	}
}

func TestBinarySearchIteration(t *testing.T) {
	for i := 0; i < testTimes; i++ {
		xs := genRandomIntsNoEmpty(testScaleLimit)
		sort.Ints(xs)
		targetIdx := rand.Intn(len(xs))
		target := xs[targetIdx]
		idx, isFind := binarySearchIteration(xs, target)
		if isFind != true || idx != targetIdx {
			fmt.Printf("target: xs[%d]=%d  found: xs[%d]=%d\n", targetIdx, target, idx, xs[idx])
			t.FailNow()
		}
	}
}
