package main

import (
	lib_heap "container/heap"
	"fmt"
)

func main() {
	// fmt.Println(lecture_sorts.SelectionSort([]int{1, 3, 0, -1, 2}))

	// fmt.Println(lecture_sorts.InsertionSort([]int{1, 3, 0, -1, 2}))

	// fmt.Println(lecture_sorts.BubbleSort([]int{1, 3, 0, -1, 2}))

	// fmt.Println(lecture_sorts.MergeSort([]int{1, 3, 0, -1, 2}))

	// fmt.Println(sem_4.ShellSort([]int{1, 3, 0, -1, 2}))

	// sem_4.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	// sem_5.PrintTreePreOrder(sem_5.BuildTree([]int{8, 9, 11, 7, 16, 3, 1}, 0))
	// 	ints := []int{1, 2}
	// 	for kek := range ints {
	// 		ints = append(ints, kek)
	// 		fmt.Println(ints)

	// 	}
	// 	fmt.Println("END", ints)
	// }
	// tree := sem_5.BuildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)

	// bft(tree)

	// var kek []int
	// fmt.Println(append(kek, 1))

	// a := 1
	// b := 2

	// a, b = b, a

	// fmt.Println(a, b)
	// k := map[string]kek{}
	// fmt.Println(k)
	q := &heap{
		heapVertex{"a", 1},
		heapVertex{"a", 4},
		heapVertex{"a", 2},
		heapVertex{"a", 0}}

	lib_heap.Init(q)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

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

// func bft(node *sem_5.Node) {

// 	queue := []*sem_5.Node{node}

// 	for len(queue) > 0 {

// 		newQueue := []*sem_5.Node{}

// 		for _, node := range queue {
// 			fmt.Println(node.Val)
// 			if node.Left != nil {
// 				newQueue = append(newQueue, node.Left)
// 			}
// 			if node.Right != nil {
// 				newQueue = append(newQueue, node.Right)
// 			}
// 		}

// 		queue = newQueue
// 	}

// }
