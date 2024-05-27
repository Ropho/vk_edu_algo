package sem_6

import (
	"fmt"

	"vk_algo/sem_5"
)

type dfsResult struct {
	CurMaxSum  int
	CurMaxPath []int

	MaxSum  int
	MaxPath []int
}

func dfs(node *sem_5.Node) *dfsResult {

	res := &dfsResult{}

	if node == nil {
		return &dfsResult{}
	}

	resLeft := dfs(node.Left)
	resRight := dfs(node.Right)

	if resLeft.CurMaxSum > resRight.CurMaxSum {
		res.CurMaxPath = append(resLeft.CurMaxPath, node.Val)
	} else {
		res.CurMaxPath = append(resRight.CurMaxPath, node.Val)
	}

	if resLeft.CurMaxSum > resRight.CurMaxSum {
		res.CurMaxSum = resLeft.CurMaxSum
	} else {
		res.CurMaxSum = resRight.CurMaxSum
	}
	res.CurMaxSum += node.Val

	checkMaxSum := resLeft.CurMaxSum + node.Val + resRight.CurMaxSum
	if checkMaxSum > resLeft.MaxSum && checkMaxSum > resRight.MaxSum {
		res.MaxSum = resLeft.CurMaxSum + node.Val + resRight.CurMaxSum

		res.MaxPath = append(resLeft.CurMaxPath, node.Val)
		res.MaxPath = append(res.MaxPath, resRight.CurMaxPath...)

	} else if resLeft.MaxSum > resRight.MaxSum {
		res.MaxSum = resLeft.MaxSum
		res.MaxPath = resLeft.MaxPath
	} else {
		res.MaxSum = resRight.MaxSum
		res.MaxPath = resRight.MaxPath
	}

	return res
}

func FindMaxPathSum(root *sem_5.Node) []int {

	result := dfs(root)

	fmt.Println(result.MaxPath)

	return result.MaxPath
}
