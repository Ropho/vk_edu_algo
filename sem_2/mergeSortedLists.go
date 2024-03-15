package sem_2

func MergeSortedLists(head1 *Node, head2 *Node) *Node {

	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head *Node
	var cur *Node

	iter1 := head1
	iter2 := head2

	if head1.Val <= head2.Val {
		head = head1
		iter1 = head1.Next
	} else {
		head = head2
		iter2 = head2.Next
	}
	cur = head

	for iter1 != nil && iter2 != nil {

		if iter1.Val <= iter2.Val {
			cur.Next = iter1
			iter1 = iter1.Next
		} else {
			cur.Next = iter2
			iter2 = iter2.Next
		}

		cur = cur.Next
	}

	if iter1 == nil {
		cur.Next = iter2
	} else {
		cur.Next = iter1
	}

	return head
}
