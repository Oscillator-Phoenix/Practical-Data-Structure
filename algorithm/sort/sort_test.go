package mysort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	testTimes      = 50
	testScaleLimit = 1000
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

func TestMergeSort(t *testing.T) {
	if sortCheck(mergeSort) == false {
		t.Fail()
	}
}

func TestQuickSort(t *testing.T) {
	if sortCheck(quickSort) == false {
		t.Fail()
	}
}

func TestInsertSort(t *testing.T) {
	if sortCheck(insertSort) == false {
		t.Fail()
	}
}

func TestSelectSort(t *testing.T) {
	if sortCheck(selectSort) == false {
		t.Fail()
	}
}
