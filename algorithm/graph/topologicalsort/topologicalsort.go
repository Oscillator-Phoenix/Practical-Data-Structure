package topologicalsort

// adjacency table
type graph struct {
	vertexs   []int
	adjacency [][]int
}

// graph should be a DAG
func topSortRecur(g *graph) []int {
	stack := []int{}

	visited := make([]bool, len(g.vertexs))
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}

	// dfs until the our-degree of node is zero
	// push the zero-outdegree-node to stack and mark that it have been visited
	var topSortHelper func(v int)
	topSortHelper = func(v int) {
		visited[v] = true
		for i := 0; i < len(g.adjacency[v]); i++ {
			to := g.adjacency[v][i]
			if visited[to] == false {
				topSortHelper(to)
			}
		}
		stack = append(stack, v)
	}

	for i := 0; i < len(g.vertexs); i++ {
		if visited[i] == false {
			topSortHelper(g.vertexs[i])
		}
	}

	// note that this recursion version cannot detect whether the graph is a DAG
	return stack
}

func topSortIter(g *graph) []int {

	topSeq := []int{}
	q := []int{}
	inDegree := make([]int, len(g.vertexs)) // zero initialize

	for from := 0; from < len(g.vertexs); from++ {
		for to := 0; to < len(g.adjacency[from]); to++ {
			inDegree[g.adjacency[from][to]]++
		}
	}

	// push the zero-in-degree node to queue
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			q = append(q, i) // queue push
		}
	}

	for len(q) > 0 {
		zeroInDegreeVertex := q[0]
		q = q[1:] // queue pop
		topSeq = append(topSeq, zeroInDegreeVertex)

		for _, to := range g.adjacency[zeroInDegreeVertex] {

			// remove the zero-in-degree node from the graph and reduce the in-degree
			inDegree[to]--

			if inDegree[to] == 0 {
				q = append(q, to)
			}
		}
	}

	if len(topSeq) != len(g.vertexs) {
		return []int{} // the graph is not a DAG and dosen't have topological sequence
	}

	return topSeq
}
