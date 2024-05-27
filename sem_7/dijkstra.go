package sem7

import (
	"fmt"

	lib_heap "container/heap"
)

const inf = 666

func Dijkstra(graph map[string]map[string]int, start string) map[string]int {

	dist := map[string]int{}
	for vertex := range graph {
		dist[vertex] = inf
	}
	dist[start] = 0

	pQueue := heap{heapVertex{
		Name: start,
		Dist: 0,
	}}
	lib_heap.Init(&pQueue)

	visited := map[string]struct{}{}

	for pQueue.Len() > 0 {
		fmt.Println("QUEUE", pQueue)

		v, ok := pQueue.Pop().(heapVertex)
		if !ok {
			return nil
		}
		fmt.Println(v)

		// lazy cleaning
		_, ok = visited[v.Name]
		if ok {
			continue
		}

		visited[v.Name] = struct{}{}

		neighbors, ok := graph[v.Name]
		if !ok {
			return nil
		}
		for name, val := range neighbors {

			if dist[name] > dist[v.Name]+val {
				dist[name] = dist[v.Name] + val
			}

			if _, ok := visited[name]; !ok {
				pQueue.Push(heapVertex{Name: name, Dist: dist[name]})
			}
		}
		lib_heap.Init(&pQueue)

	}

	return dist
}

type heapVertex struct {
	Name string
	Dist int
}

type heap []heapVertex

func (h heap) Len() int {
	return len(h)
}

func (h heap) Less(i int, j int) bool {
	return h[i].Dist > h[j].Dist
}

func (h heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heap) Push(x any) {
	*h = append(*h, x.(heapVertex))
}

func (h *heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
