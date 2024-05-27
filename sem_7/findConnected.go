package sem7

func dfs(graph map[int][]int, v int, visited map[int]int, color int) {
	visited[v] = color

	for _, i := range graph[v] {
		if visited[i] == 0 {
			dfs(graph, i, visited, color)
		}
	}
}

func findConnectedComponents(graph map[int][]int) [][]int {
	visited := map[int]int{}

	color := 0
	for name := range graph {
		visited[name] = color
	}

	for node := range graph {
		if visited[node] == 0 {
			color++
			dfs(graph, node, visited, color)
		}
	}

	ans := make([][]int, color)
	for name, color := range visited {
		ans[color-1] = append(ans[color-1], name)
	}

	return ans
}
