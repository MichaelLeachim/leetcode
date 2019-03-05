// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-04 20:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreePostorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	assert.Equal(t, postorderTraversalIterative(tree), []int{3, 2, 1})
	assert.Equal(t, postOrderTraversalRecursive(tree), []int{3, 2, 1})

	tree = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	assert.Equal(t, postorderTraversalIterative(tree), []int{1, 2, 3})
	assert.Equal(t, postOrderTraversalRecursive(tree), []int{1, 2, 3})
	assert.Equal(t, postorderTraversalIterative(nil), []int{})

}
