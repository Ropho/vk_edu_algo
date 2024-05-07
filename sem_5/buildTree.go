package sem_5

func BuildTree(arr []int, index int) *Node {
	if arr[index] == Poison {
		return nil
	}

	node := &Node{
		Val: arr[index],
	}

	if 2*index+1 < len(arr) {
		node.Left = BuildTree(arr, 2*index+1)
	}

	if 2*index+2 < len(arr) {
		node.Right = BuildTree(arr, 2*index+2)
	}

	return node
}
