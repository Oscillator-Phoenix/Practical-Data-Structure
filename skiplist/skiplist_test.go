package skiplist

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestRandomLevel(t *testing.T) {
	// todo
}

func TestInsert1(t *testing.T) {
	sl := newSkipList()
	// fmt.Println(sl)

	xs := []int{3, 6, 7, 9, 12, 17, 19, 21, 25, 26}
	rand.Shuffle(len(xs), func(i, j int) {
		xs[i], xs[j] = xs[j], xs[i]
	})
	// fmt.Println("Shuffled xs: ", xs)

	for i := 0; i < len(xs); i++ {
		if err := sl.insert(xs[i], xs[i]); err != nil {
			t.Logf("faild at insert element xs[%d]=%d", i, xs[i])
			t.Log("xs", xs)
			t.Fail()
		}
	}

	ys := []int{}
	sl.traverse(func(key, value int) {
		// fmt.Printf("(%d, %d) ", key, value)
		ys = append(ys, value)
	})

	sort.Ints(xs)
	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			t.Fail()
		}
	}
}

func TestInsert2(t *testing.T) {
	build := func(xs []int, sl *skipList) {
		for i := 0; i < len(xs); i++ {
			if err := sl.insert(xs[i], xs[i]); err != nil {
				t.Logf("faild at insert element xs[%d]=%d", i, xs[i])
				t.Fail()
			}
		}
	}

	check := func(xs []int, sl *skipList) {
		ys := []int{}
		sl.traverse(func(key, value int) {
			ys = append(ys, value)
		})
		sort.Ints(xs)
		for i := 0; i < len(xs); i++ {
			if xs[i] != ys[i] {
				t.Fail()
			}
		}
	}

	const (
		testTimes      = 100
		testScaleLimit = 10000
		testNumRange   = 100000
	)

	newRandomInts := func(size int) []int {
		xs := map[int](struct{}){}
		for i := 0; i < size; i++ {
			xs[rand.Intn(testNumRange)] = struct{}{}
		}
		_xs := []int{}
		for x := range xs {
			_xs = append(_xs, x)
		}
		return _xs
	}

	newRandomScale := func(size int) []int {
		randtScale := newRandomInts(size)
		for i := 0; i < len(randtScale); i++ {
			randtScale[i] = randtScale[i] % testScaleLimit
		}
		return randtScale
	}

	scales := newRandomScale(testTimes)
	for i := 0; i < len(scales); i++ {
		xs := newRandomInts(scales[i])
		sl := newSkipList()
		build(xs, &sl)
		check(xs, &sl)

		fmt.Printf("tests %d / %d (size: %d) passed\n", i+1, len(scales), len(xs))

	}
}
