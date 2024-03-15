package sem_2

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Note struct {
	Val       int
	NextIndex int
}

func IsEqual(head_1 *Node, head_2 *Node) bool {

	for head_1 != nil && head_2 != nil {
		if head_1.Val != head_2.Val {
			return false
		}

		head_1 = head_1.Next
		head_2 = head_2.Next
	}

	if head_1 == nil && head_2 == nil {
		return true
	}

	return false
}

func PrintAll(head *Node) {

	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}

	fmt.Print("END\n")
}

// [ {1, 1}, {2, -1}]
func ConstructList(notes []Note) *Node {

	if len(notes) == 0 {
		return nil
	}

	var head *Node
	nodeAddr := make([]*Node, len(notes))

	for index := range notes {

		nodeAddr[index] = &Node{
			Val:  notes[index].Val,
			Next: nil,
		}
	}

	for index := range notes {

		nextIndex := notes[index].NextIndex
		if nextIndex != -1 {
			nodeAddr[index].Next = nodeAddr[notes[index].NextIndex]
		}
	}

	head = nodeAddr[0]

	return head
}
