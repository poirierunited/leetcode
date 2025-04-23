package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func main() {
	root := &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 1,
			right: &TreeNode{
				value: 3,
			},
		},
		right: &TreeNode{
			value: 5,
			right: &TreeNode{
				value: 1,
			},
		},
	}

	withRoot, withoutRoot := dfs(root)
	result := max(withRoot, withoutRoot)
	fmt.Printf("Maximum amount that can be robbed: %d\n", result)
}

func dfs(node *TreeNode) (int, int) {
	// base case
	if node == nil {
		return 0, 0
	}

	leftWith, leftWithout := dfs(node.left)
	rightWith, rightWithout := dfs(node.right)

	withRoot := node.value + leftWithout + rightWithout
	withoutRoot := max(leftWith, leftWithout) + max(rightWith, rightWithout)

	return withRoot, withoutRoot
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
