package sem7

import "fmt"

func IsTree(graph map[int][]int, start int) bool {
	visited := map[int]struct{}{}
	queue := []int{start}
	parent := map[int]int{start: -1}

	for len(queue) > 0 {
		fmt.Println(queue)

		vertex := queue[0]
		queue = queue[1:]
		visited[vertex] = struct{}{}

		for _, neighbor := range graph[vertex] {

			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)

				parent[neighbor] = vertex
			} else {
				if parent[vertex] != neighbor {
					return false
				}
			}
		}
	}

	return len(visited) == len(graph)
}
