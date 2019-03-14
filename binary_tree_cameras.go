// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-13 12:11 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/binary-tree-cameras/

// When we **have to** add camera to a node?
//   if parent has camera, then none

//   when parent have no camera and none of its children have camera
//   when none of its children exist:
//     if parent has camera, then none, otherwise yes
//
//   when one of its children does not exist, then, the reverse
//   add camera.

// When none of the surroundings has camera
// When there are no surroundings, then true

type binary_tree_cameras_ struct {
	_memoized map[*TreeNode]bool
}

var binary_tree_cameras = binary_tree_cameras_{
	_memoized: map[*TreeNode]bool{},
}

func (b *binary_tree_cameras_) Solve(t *TreeNode) bool {
	res, ok := b._memoized[t]
	if ok {
		return res
	}

	if t.Left == nil && t.Right == nil {
		b._memoized[t] = false
		return false
	}
	if t.Left == nil {
		res := !b.Solve(t.Right)
		b._memoized[t] = res
		return res
	}
	if t.Right == nil {
		res := !b.Solve(t.Left)
		b._memoized[t] = res
		return res
	}

	if b.Solve(t.Left) == false && b.Solve(t.Right) == false {
		b._memoized[t] = true
		return true
	}
	b._memoized[t] = false
	return false

}

func (b *binary_tree_cameras_) CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + b.CountNodes(root.Left) + b.CountNodes(root.Right)
}

func minCameraCover(root *TreeNode) int {

	binary_tree_cameras._memoized = map[*TreeNode]bool{}
	binary_tree_cameras.Solve(root)

	// patch those nodes where solution didn't get to
	// (have no idea, whether it will work, though)
	count := 0
	for _, a := range binary_tree_cameras._memoized {
		if a == true {
			count += 1
		}

	}
	return count
}
