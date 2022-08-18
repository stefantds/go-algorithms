package graphs

func BFS(n int, adj [][]int, s, t int) (bool, []int) {
	pathTo := make([]int, n)
	for i := range pathTo {
		pathTo[i] = -1
	}
	marked := make([]bool, n)
	queue := make([]int, 0)

	queue = append(queue, s)
	marked[s] = true
	found := false
	for len(queue) > 0 && !found {
		var cur int
		cur, queue = queue[0], queue[1:]

		for _, n := range adj[cur] {
			if !marked[n] {
				marked[n] = true
				pathTo[n] = cur
				if n == t {
					found = true
					break
				}
				queue = append(queue, n)
			}
		}
	}

	if !found {
		return false, []int{}
	}

	path := make([]int, 0)
	for i := t; i != s && i != -1; i = pathTo[i] {
		path = append(path, i)
	}
	path = append(path, s)

	return true, reverse(path)
}
