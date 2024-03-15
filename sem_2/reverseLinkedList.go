package sem_2

func ReverseLinkedList(head *Node) *Node {

	var prev *Node
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	head = prev
	return head
}
