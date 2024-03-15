package sem_2

func DeleteNode(head *Node, val int) *Node {

	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	cur := head.Next
	prev := head
	for cur != nil {

		if cur.Val == val {
			prev.Next = cur.Next
			break
		}
		cur = cur.Next
		prev = prev.Next
	}

	return head
}
