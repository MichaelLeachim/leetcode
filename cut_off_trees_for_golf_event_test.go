// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-01 12:02 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceGet(t *testing.T) {
	assert.Equal(t, canSlice(len([]int{}), 1), false)
	assert.Equal(t, canSlice(len([]int{}), 0), false)
	assert.Equal(t, canSlice(len([]int{}), -1), false)
	assert.Equal(t, canSlice(len([]int{1}), 0), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 0), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 5), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 6), false)
	for index, item := range []int{1, 2, 3, 4, 5, 6} {
		if canSlice(len([]int{1, 2, 3, 4, 5, 6}), index) {
			continue
		}
		assert.Fail(t, "this branch should be unreachable")
	}

}
func TestConstructAdjacencyMatrixRepresentation(t *testing.T) {

}
