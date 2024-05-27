package sem7

func hasCycle(graph map[int][]int) bool {
	visited := map[int]bool{}

	for vertex := range graph {
		if !visited[vertex] {
			if dfsCycle(graph, vertex, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfsCycle(graph map[int][]int, vertex int, parent int, visited map[int]bool) bool {
	visited[vertex] = true

	for _, neighbor := range graph[vertex] {
		if neighbor != parent {
			if visited[neighbor] || dfsCycle(graph, neighbor, vertex, visited) {
				return true
			}
		}
	}
	return false
}
