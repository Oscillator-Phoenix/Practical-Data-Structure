package mysort

// recursion! recursion! recursion!

type lessFunc func(x, y int) bool

type sortFunc func(xs []int, less lessFunc) []int

func merge(xs, ys []int, less lessFunc) []int {
	i := 0
	j := 0
	zs := []int{}

	for i < len(xs) && j < len(ys) {
		if less(xs[i], ys[j]) {
			zs = append(zs, xs[i])
			i++
		} else {
			zs = append(zs, ys[j])
			j++
		}
	}

	for i < len(xs) {
		zs = append(zs, xs[i])
		i++
	}

	for j < len(ys) {
		zs = append(zs, ys[j])
		j++
	}

	return zs
}

func mergeSort(xs []int, less lessFunc) []int {
	if xs == nil {
		return nil
	}

	if len(xs) == 0 || len(xs) == 1 {
		return xs
	}

	mid := (len(xs) + 1) / 2
	left := mergeSort(xs[:mid], less)
	right := mergeSort(xs[mid:], less)

	return merge(left, right, less)
}

func quickSort(xs []int, less lessFunc) []int {
	if xs == nil {
		return nil
	}
	if len(xs) == 0 {
		return []int{}
	}
	if len(xs) == 1 {
		return []int{xs[0]}
	}

	pivot := xs[0]
	left := []int{}
	right := []int{}
	for i := 1; i < len(xs); i++ {
		if less(xs[i], pivot) {
			left = append(left, xs[i])
		} else {
			right = append(right, xs[i])
		}
	}

	sorted := []int{}
	sorted = append(sorted, quickSort(left, less)...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, quickSort(right, less)...)

	return sorted
}

func insertToSorted(y int, xs []int, less lessFunc) []int {
	pos := 0
	for pos < len(xs) {
		if less(xs[pos], y) {
			pos++
		} else {
			break
		}
	}

	ys := []int{}
	ys = append(ys, xs[:pos]...)
	ys = append(ys, y)
	ys = append(ys, xs[pos:]...)

	return ys
}

func insertSort(xs []int, less lessFunc) []int {
	if xs == nil {
		return nil
	}
	if len(xs) == 0 {
		return []int{}
	}
	return insertToSorted(xs[0], insertSort(xs[1:], less), less)
}

// selectMin requires: len(xs) >= 1
func selectMin(xs []int, less lessFunc) (int, []int) {
	if len(xs) == 1 {
		return xs[0], []int{}
	}

	minIdx := 0
	min := xs[0]
	for i := 0; i < len(xs); i++ {
		if xs[i] < min {
			minIdx = i
			min = xs[i]
		}
	}
	remain := []int{}
	remain = append(remain, xs[:minIdx]...)
	remain = append(remain, xs[(minIdx+1):]...)

	return min, remain
}

func selectSort(xs []int, less lessFunc) []int {
	if xs == nil {
		return nil
	}
	if len(xs) == 0 {
		return []int{}
	}

	min, remain := selectMin(xs, less)

	sorted := []int{}
	sorted = append(sorted, min)
	sorted = append(sorted, selectSort(remain, less)...)

	return sorted
}

func bubbleSort(xs []int, less lessFunc) []int {
	return nil
}
