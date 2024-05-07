package sem_5

func IsSymmetricBFS(root *Node) bool {

	if root == nil {
		return false
	}

	queue := []*Node{root}
	for len(queue) != 0 {

		for index := 0; index < len(queue)/2; index++ {
			if queue[index].Val != queue[len(queue)-1-index].Val {
				return false
			}
		}

		newQueue := []*Node{}

		for _, node := range queue {

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}

		}
		queue = newQueue
	}

	return true
}

// IN-ORDER
func IsSymmetricDFS(root *Node) bool {
	if root == nil {
		return false
	}

	inOrderNodes := []*Node{}

	inOrder(root, &inOrderNodes)

	for index := 0; index < len(inOrderNodes)/2; index++ {
		if inOrderNodes[index].Val != inOrderNodes[len(inOrderNodes)-1-index].Val {
			return false
		}
	}

	return true
}
