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

	// check memoization
	a := ntn(0, nil, nil)
	binary_tree_cameras.memoize(a, true, 12, false)
	ok, _, _ := binary_tree_cameras.unmemoize(a, false)
	assert.Equal(t, ok, false)
	ok, count, isChild := binary_tree_cameras.unmemoize(a, true)
	assert.Equal(t, ok, true)
	assert.Equal(t, count, 12)
	assert.Equal(t, isChild, false)

	bt := ntn(0, ntn(0, ntn(0, ntn(0, nil, ntn(0, nil, nil)), nil), nil), nil)
	bt2 := ntn(0, ntn(0, ntn(0, nil, nil), ntn(0, nil, nil)), nil)
	assert.Equal(t, 2, minCameraCover(bt))
	assert.Equal(t, 1, minCameraCover(bt2))
	assert.Equal(t, minCameraCover(ntn(0, nil, nil)), 1)

	bt4 := ntn(0, ntn(0, nil, nil), ntn(0, nil, ntn(0, nil, nil)))
	assert.Equal(t, minCameraCover(bt4), 2)

	assert.Equal(t, 2, minCameraCover(
		ntn(0, nil,
			ntn(0,
				ntn(0, nil, nil),
				ntn(0, nil,
					ntn(0, nil, nil),
				),
			),
		),
	))

	// // [0,null,0,0,0,null,null,null,0]
	// // bt5 := ntn(0,nil,ntn(0,ntn(0,)
}

func BenchmarkMinCameraCover(t *testing.B) {
	ntn := func(v int, l, r *TreeNode) *TreeNode {
		return &TreeNode{Val: v, Left: l, Right: r}
	}
	bt := ntn(0, ntn(0, ntn(0, ntn(0, nil, ntn(0, nil, nil)), nil), nil), nil)
	for n := 0; n < t.N; n++ {
		minCameraCover(bt)
	}
}
