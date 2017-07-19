package oj

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type OjError struct {
	Msg string
}

func NewTree(values []int) (top *TreeNode, err error) {
	if len(values) > 0 {

	} else {
		err = OjError{
			"zero value array",
		}
	}
	return
}
