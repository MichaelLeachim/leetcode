// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-04 20:16 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

// postorder traversal. first, traverse left subtree, then, traverse right subtree,then, process the root

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const isIterative = true

// working iterative solution
func postorderTraversalIterative(rootNode *TreeNode) []int {
	result := []int{}

	afterTraverseRight := func(root *TreeNode) {
		result = append(result, root.Val)
	}
	afterTraverseLeft := func(stack []*TreeNode, root *TreeNode) {
		if root.Right != nil {
			stack = append(stack, root.Right)
			return
		}
		afterTraverseRight(root)
	}

	type StackFrame struct {
		t *TreeNode
		p int
	}
	var root *TreeNode
	var stackFrame *StackFrame

	stack := []*StackFrame{&StackFrame{t: rootNode, p: 1}}
	for len(stack) != 0 {
		stack, stackFrame = stack[:len(stack)-1], stack[len(stack)-1]
		root = stackFrame.t
		// implementation of jumps in stackframe
		switch stackFrame.p {
		default:
			if root.Left != nil {

				continue
			}
			if root.Right != nil {
				stack = append(stack, &StackFrame{p: 4, t: root.Right})
				continue
			}
			result = append(result, root.Val)
		case 1:
			if root.Left != nil {
				stack = append(stack, &StackFrame{p: 3, t: root.Left})
				continue
			}
		case 2:
			if root.Right != nil {
				stack = append(stack, &StackFrame{p: 4, t: root.Right})
				continue
			}
		}
		return result
	}
	return []int{}
}

// working recursive solution
func postOrderTraversalRecursive(root *TreeNode) []int {
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

func postorderTraversal(root *TreeNode) []int {
	if isIterative {
		return postorderTraversalIterative(root)
	}
	return postOrderTraversalRecursive(root)

}
