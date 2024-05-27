package sem7

func BellmanFord(graph map[string]map[string]int, start string) map[string]int {

	dist := map[string]int{}
	for v := range graph {
		dist[v] = inf
	}
	dist[start] = 0

	for i := 0; i < len(graph)-1; i++ {
		for from, neighbors := range graph {

			for to, curDist := range neighbors {

				if dist[from]+curDist < dist[to] {
					dist[to] = dist[from] + curDist
				}
			}
		}
	}

	for i := 0; i < len(graph)-1; i++ {
		for from, neighbors := range graph {

			for to, curDist := range neighbors {

				if dist[from]+curDist < dist[to] {
					dist[to] = -inf
				}
			}
		}
	}

	return dist
}
