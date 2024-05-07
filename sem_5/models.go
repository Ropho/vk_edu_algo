package sem_5

import "fmt"

const Poison = 666

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func PrintTreePreOrder(node *Node) {

	fmt.Println(node.Val)

	if node.Left != nil {
		PrintTreePreOrder(node.Left)
	}
	if node.Right != nil {
		PrintTreePreOrder(node.Right)
	}

}

func inOrder(node *Node, sl *[]*Node) {

	if node.Left != nil {
		inOrder(node.Left, sl)
	}

	*sl = append(*sl, node)

	if node.Right != nil {
		inOrder(node.Right, sl)
	}
}
