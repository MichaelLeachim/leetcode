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

func (b *binary_tree_cameras_) Solve(p, t *TreeNode) bool {
	res, ok := b._memoized[t]
	if ok {
		return res
	}

	if t.Left == nil && t.Right == nil {
		b._memoized[t] = false
		return false
	}
	if t.Left == nil {
		res := !b.Solve(t, t.Right)
		b._memoized[t] = res
		return res
	}
	if t.Right == nil {
		res := !b.Solve(t, t.Left)
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

func (b *binary_tree_cameras_) BuildBinaryTreeFromSlice(k int, input []int) *TreeNode {
	l, r := 2*k, 2*k+1
	if k < len(input) {
		if input[k] == -1 {
			return nil
		}
		return &TreeNode{
			Val:   input[k],
			Left:  b.BuildBinaryTreeFromSlice(l, input),
			Right: b.BuildBinaryTreeFromSlice(r, input),
		}
	}
	return nil
}
func (b *binary_tree_cameras_) ShowTree(tree *TreeNode, logger func(args ...interface{})) {

	var showrecur func(tree *TreeNode, pad string, name string, logger func(args ...interface{}))
	showrecur = func(tree *TreeNode, pad string, name string, logger func(args ...interface{})) {
		if tree != nil {
			logger(pad, name, tree.Val)
			showrecur(tree.Left, pad+"-", "[l]", logger)
			showrecur(tree.Right, pad+"-", "[r]", logger)
		}
	}
	showrecur(tree, "-", "", logger)
}

func (b *binary_tree_cameras_) CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + b.CountNodes(root.Left) + b.CountNodes(root.Right)
}

func (b *binary_tree_cameras_) BuildSliceFromBinaryTreeRecursive(data []int, k int, root *TreeNode) {
	if root == nil {
		data[k] = -1
		return
	}
	if k < len(data) {
		data[k] = root.Val
		b.BuildSliceFromBinaryTreeRecursive(data, k*2+1, root.Left)
		b.BuildSliceFromBinaryTreeRecursive(data, k*2+2, root.Right)
	}

}

func (b *binary_tree_cameras_) BuildSliceFromBinaryTree(root *TreeNode) []int {
	result := make([]int, b.CountNodes(root))
	b.BuildSliceFromBinaryTreeRecursive(result, 0, root)
	return result
}

func minCameraCover(root *TreeNode) int {
	binary_tree_cameras._memoized = map[*TreeNode]bool{}
	// special case that I haven't thought about
	if root.Left == nil && root.Right == nil {
		return 1
	}

	binary_tree_cameras.Solve(root)
	count := 0
	for _, a := range binary_tree_cameras._memoized {
		if a == true {
			count += 1
		}

	}
	return count
}
