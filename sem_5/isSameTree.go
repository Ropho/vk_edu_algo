package sem_5

func IsSameTree(first *Node, second *Node) bool {

	if first == nil && second == nil {
		return true
	}

	if first == nil || second == nil {
		return false
	}

	if first.Val != second.Val {
		return false
	}

	left := IsSameTree(first.Left, second.Left)
	right := IsSameTree(first.Right, second.Right)

	return left && right
}
