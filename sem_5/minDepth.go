package sem_5

func MinDepth(root *Node) int {

	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {

		left := MinDepth(root.Left)
		right := MinDepth(root.Right)
		if left <= right {
			return 1 + left
		}

		return 1 + right
	}

	if root.Left != nil {
		return 1 + MinDepth(root.Left)
	}

	if root.Right != nil {
		return 1 + MinDepth(root.Right)
	}

	return 1
}
