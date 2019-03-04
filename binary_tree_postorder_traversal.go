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

// postorder traversal. first, traverse left subtree, then, traverse right subtree,then, process the root
func postorderTraversal(root *TreeNode) []int {
	var traverse func(root *TreeNode)
	result := []int{}
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
