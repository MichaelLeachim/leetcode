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
	if rootNode == nil {
		return []int{}
	}
	type StackFrame struct {
		t *TreeNode
		p int
	}
	root := rootNode
	var stackFrame *StackFrame

	result := []int{}

	stack := []*StackFrame{&StackFrame{t: rootNode, p: 1}}

	for len(stack) != 0 {
		if root == nil {
			// exit operation. Pop out stack frame, change funtion context
			stack, stackFrame = stack[:len(stack)-1], stack[len(stack)-1]
			root = stackFrame.t
			switch stackFrame.p {
			case 2:
				goto second
			case 3:
				goto third
			}
		}
		if root != nil && root.Left != nil {
			// save current execution context onto stack
			stack = append(stack, &StackFrame{p: 2, t: root})
			// change current execution context into recurrent one
			root = root.Left
			continue
		}
	second:
		if root != nil && root.Right != nil {
			// save current execution context onto stack
			stack = append(stack, &StackFrame{p: 3, t: root})
			// change current execution context into recurrent one
			root = root.Right
			continue
		}
	third:
		result = append(result, root.Val)
		// return from recursion:
		stack, stackFrame = stack[:len(stack)-1], stack[len(stack)-1]
		// restore parent execution context
		root = stackFrame.t
		switch stackFrame.p {
		case 2:
			goto second
		case 3:
			goto third
		}
	}
	return result
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
		return
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
