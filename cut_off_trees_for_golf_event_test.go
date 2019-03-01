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
	_, ok := sliceGet([]int{}, 1)
	assert.Equal(t, ok, false)

	_, ok = sliceGet([]int{}, 0)
	assert.Equal(t, ok, false)

	_, ok = sliceGet([]int{}, -1)
	assert.Equal(t, ok, false)

	item, ok := sliceGet([]int{1}, 0)
	assert.Equal(t, ok, true)
	assert.Equal(t, item, 1)

	item, ok = sliceGet([]int{1, 2, 3, 4, 5, 6}, 0)
	assert.Equal(t, ok, true)
	assert.Equal(t, item, 1)

	item, ok = sliceGet([]int{1, 2, 3, 4, 5, 6}, 5)
	assert.Equal(t, ok, true)
	assert.Equal(t, item, 6)

	_, ok = sliceGet([]int{1, 2, 3, 4, 5, 6}, 6)
	assert.Equal(t, ok, false)

}
func TestConstructAdjacencyMatrixRepresentation(t *testing.T) {

}
