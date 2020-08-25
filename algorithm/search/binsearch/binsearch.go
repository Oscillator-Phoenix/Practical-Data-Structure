package binsearch

func binarySearchIteration(xs []int, target int) (idx int, isFind bool) {
	left := 0
	right := len(xs)

	for left < right {
		mid := (left + right) / 2
		if xs[mid] < target {
			left = mid + 1
		} else if xs[mid] > target {
			right = mid - 1
		} else {
			return mid, true
		}
	}

	return -1, false
}

func binarySearchRecursion(xs []int, target int) (idx int, isFind bool) {
	if len(xs) == 0 {
		return -1, false
	}

	mid := len(xs) / 2

	if xs[mid] > target {
		if idx, isFind = binarySearchRecursion(xs[:mid], target); isFind {
			return idx, isFind
		}
	} else if xs[mid] < target {
		if idx, isFind = binarySearchRecursion(xs[(mid+1):], target); isFind {
			return mid + 1 + idx, isFind
		}
	}

	return mid, true
}
