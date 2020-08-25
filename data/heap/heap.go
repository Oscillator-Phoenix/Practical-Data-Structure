package heap

type heap struct {
	xs []int
}

// min heap
func newHeap() heap {
	return heap{
		xs: make([]int, 0),
	}
}

func (hp *heap) size() int {
	return len(hp.xs)
}

func (hp *heap) empty() bool {
	return len(hp.xs) == 0
}

func (hp *heap) top() int {
	if hp.empty() {
		panic("heap is empty")
	}
	return hp.xs[0]
}

func (hp *heap) push(x int) {
	hp.xs = append(hp.xs, x)
	hp.shiftUp()
}

func (hp *heap) pop() int {
	top := hp.top()

	hp.xs[0] = hp.xs[len(hp.xs)-1]
	hp.xs = hp.xs[:len(hp.xs)-1]
	hp.shiftDown()

	return top
}

func (hp *heap) shiftDown() {
	// fmt.Println("shift donw from: ", hp.xs)
	i := 0
	for {
		down := i
		left := 2*i + 1
		right := 2*i + 2

		if left < len(hp.xs) && hp.xs[left] < hp.xs[down] {
			down = left
		}
		if right < len(hp.xs) && hp.xs[right] < hp.xs[down] {
			down = right
		}

		if down == i {
			// fmt.Println("shift done to ", hp.xs)
			break // shift down done
		} else if down == left {
			// fmt.Println("shift ", i, " --> ", left)
			hp.xs[i], hp.xs[left] = hp.xs[left], hp.xs[i] // swap
			i = left
		} else if down == right {
			// fmt.Println("shift ", i, " --> ", right)
			hp.xs[i], hp.xs[right] = hp.xs[right], hp.xs[i] // swap
			i = right
		}
	}
}

func (hp *heap) shiftUp() {
	i := len(hp.xs) - 1
	for {
		up := i
		father := (i - 1) / 2

		if father >= 0 && hp.xs[father] > hp.xs[up] {
			up = father
		}

		if up == i {
			break // shift up done
		} else if up == father {
			hp.xs[i], hp.xs[father] = hp.xs[father], hp.xs[i]
			i = father
		}
	}
}

func heapSort(xs []int) {
	hp := newHeap()

	for i := 0; i < len(xs); i++ {
		hp.push(xs[i])
	}

	for i := 0; i < len(xs); i++ {
		xs[i] = hp.pop()
	}
}
