package heap

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestHeapPop(t *testing.T) {
	fmt.Println("TestHeapPop --------------------------------------> ")
	hp := newHeap()
	hp.xs = []int{0, 1, 2, 3, 4, 5}

	fmt.Println("size of heap: ", hp.size())

	sz := hp.size()
	for i := 0; i < sz; i++ {
		fmt.Println("pop: ", hp.pop(), " heap: ", hp.xs)
	}
}

const (
	testTimes      = 100
	testScaleLimit = 1000
	testNumRange   = 100000
)

func newRandomInts(size int) []int {
	xs := make([]int, size)
	for i := 0; i < size; i++ {
		xs[i] = rand.Intn(10000)
	}
	return xs
}

func newRandomScale(size int) []int {
	randtScale := newRandomInts(size)
	for i := 0; i < len(randtScale); i++ {
		randtScale[i] = randtScale[i] % testScaleLimit
	}
	return randtScale
}

func TestHeapSort(t *testing.T) {
	rand.Seed(42)
	scales := newRandomScale(testTimes)
	// fmt.Println("scales: ", scales)

	tests := make([][]int, len(scales))
	for i := 0; i < len(tests); i++ {
		tests[i] = newRandomInts(scales[i])
	}
	// fmt.Println("tests: ", tests)

	for i := 0; i < len(tests); i++ {
		xs := tests[i]
		xsCopy := make([]int, len(xs))
		copy(xsCopy, xs)

		// fmt.Println("xs", xs)
		// fmt.Println("xsCopy", xsCopy)

		heapSort(xs)
		sort.Ints(xsCopy)
		for i := 0; i < len(xs); i++ {
			if xs[i] != xsCopy[i] {
				t.Logf("failed at %d th test", i)
				t.Fail()
			}
		}

		// fmt.Println("sorted xs", xs)
		// fmt.Println("sorted xsCopy", xsCopy)

		fmt.Printf("tests %d / %d (size: %d) passed\n", i+1, len(tests), len(tests[i]))
	}
}
