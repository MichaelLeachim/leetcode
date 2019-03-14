// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-13 12:59 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCameraCover(t *testing.T) {
	ntn := func(v int, l, r *TreeNode) *TreeNode {
		return &TreeNode{Val: v, Left: l, Right: r}
	}
	// bts := []int{0, 0, -1, 0, -1, 0, -1, -1, -1, -1, 0}e
	bt := ntn(0, ntn(0, ntn(0, ntn(0, nil, ntn(0, nil, nil)), nil), nil), nil)
	binary_tree_cameras.ShowTree(bt, t.Log)

	assert.Equal(t, 5, binary_tree_cameras.CountNodes(bt))
	assert.Equal(t, binary_tree_cameras.Solve(bt), false)
	assert.Equal(t, binary_tree_cameras.Solve(bt.Left), true)
	assert.Equal(t, binary_tree_cameras.Solve(bt.Left.Left), false)
	assert.Equal(t, binary_tree_cameras.Solve(bt.Left.Left.Left), true)

	bt2 := ntn(0, ntn(0, ntn(0, nil, nil), ntn(0, nil, nil)), nil)
	assert.Equal(t, binary_tree_cameras.Solve(bt2), false)
	assert.Equal(t, binary_tree_cameras.Solve(bt2.Left), true)
	assert.Equal(t, binary_tree_cameras.Solve(bt2.Left.Right), false)
	assert.Equal(t, binary_tree_cameras.Solve(bt2.Left.Left), false)

	assert.Equal(t, 2, minCameraCover(bt))
	assert.Equal(t, 1, minCameraCover(bt2))

	assert.Equal(t, minCameraCover(ntn(0, nil, nil)), 1)

	bt4 := ntn(0, ntn(0, nil, nil), ntn(0, nil, ntn(0, nil, nil)))
	assert.Equal(t, minCameraCover(bt4), 2)
}

func BenchmarkMinCameraCover(t *testing.B) {
	ntn := func(v int, l, r *TreeNode) *TreeNode {
		return &TreeNode{Val: v, Left: l, Right: r}
	}
	bt := ntn(0, ntn(0, ntn(0, ntn(0, nil, ntn(0, nil, nil)), nil), nil), nil)
	for n := 0; n < t.N; n++ {
		binary_tree_cameras.Solve(bt)
	}
}
