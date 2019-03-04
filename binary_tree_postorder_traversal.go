// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-04 20:16 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const isIterative = false

// postorder traversal. first, traverse left subtree, then, traverse right subtree,then, process the root
func postorderTraversal(root *TreeNode) []int {
	if isIterative {
		var root *TreeNode
		result := []int{}
		stack := []*TreeNode{root}
		for len(stack) != 0 {
			stack, root = stack[:len(stack)-1], stack[len(stack)-1]
			if root == nil {
				continue
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
				continue
			}
			if root.Right != nil {
				stack = append(stack, root.Right)
				continue
			}
			result = append(result, root.Val)
		}
		return result
	}

	// fully working recursive solution

	if !isIterative {
		result := []int{}

		var traverse func(root *TreeNode)
		traverse = func(root *TreeNode) {
			if root == nil {
				return
			}
			if root.Left != nil {
				traverse(root.Left)
			}
			if root.Right != nil {
				traverse(root.Right)
			}
			result = append(result, root.Val)
		}
		traverse(root)
		return result
	}
	return []int{}

}
