package main

import (
	"fmt"
	"vk_algo/sem_5"
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
	tree := sem_5.BuildTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)

	bft(tree)
}

func bft(node *sem_5.Node) {

	queue := []*sem_5.Node{node}

	for len(queue) > 0 {

		newQueue := []*sem_5.Node{}

		for _, node := range queue {
			fmt.Println(node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}

		queue = newQueue
	}

}
