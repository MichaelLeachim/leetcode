// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 16:07 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBinaryCombination(t *testing.T) {

	assert.Equal(t, strconv.FormatInt(int64(1<<3), 2), "1000")
	assert.Equal(t, strconv.FormatInt(int64(1>>3), 2), "0")
	assert.Equal(t, len(strconv.FormatInt(int64((1<<31)-1), 2)), 31)
	assert.Equal(t, len(strconv.FormatInt(int64(-(1<<31)+1), 2)), 32)
	assert.Equal(t, len(strconv.FormatUint(uint64(1<<32)-1, 2)), 32)

	assert.Equal(t, 63, len(strconv.FormatInt(int64((1<<63)-1), 2)))

	tfn := func(x, y uint32) {
		combination := uint64(x)<<32 | uint64(y)
		assert.Equal(t, x, uint32(combination>>32))
		assert.Equal(t, y, uint32((combination<<32)>>32))
	}

	tfn(123, 539)
	tfn(0, 0)
	tfn((1<<32)-1, (1<<32)-1)
}

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
	assert.Equal(t, largestRectangleArea([]int{6, 5}), 10)

	assert.Equal(t, largestRectangleArea([]int{6, 1}), 6)

	assert.Equal(t, largestRectangleArea([]int{1, 2}), 2)
	assert.Equal(t, largestRectangleArea([]int{1, 3}), 3)

	assert.Equal(t, largestRectangleArea([]int{2, 3}), 4)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 0, 1), 1)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 0, 2), 2)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 0, 3), 0)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 1, 3), 3)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 1, 2), 4)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 1, 1), 2)
	assert.Equal(t, largestRectangleAreaSolution([]int{2, 3}, 1, 0), 0)

	assert.Equal(t, largestRectangleArea([]int{3, 2}), 4)

	assert.Equal(t, largestRectangleArea([]int{5, 6}), 10)
	assert.Equal(t, largestRectangleArea([]int{1, 2, 3, 4, 5, 6, 1}), 12)

	assert.Equal(t, largestRectangleArea([]int{3, 1}), 3)
	assert.Equal(t, largestRectangleArea([]int{2}), 2)
	assert.Equal(t, largestRectangleArea([]int{2, 4}), 4)
	assert.Equal(t, largestRectangleArea([]int{2, 4, 2}), 6)
	assert.Equal(t, largestRectangleArea([]int{2, 4, 2, 1, 1, 1, 1}), 7)
	assert.Equal(t, largestRectangleArea([]int{1, 1, 1, 1, 1, 1, 10}), 10)
	assert.Equal(t, largestRectangleArea([]int{1, 1, 1, 1, 1, 1, 10, 1, 1}), 10)

	assert.Equal(t, largestRectangleArea([]int{0, 0, 0, 0, 0, 0, 0, 0, 2147483647}), 2147483647)

	assert.Equal(t, largestRectangleArea([]int{0, 0, 0, 0, 0, 0, 10, 110, 2147483647}), 2147483647)
	assert.Equal(t, largestRectangleArea([]int{2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0}), 2147483647)
	larger_input := []int{}
	for i := 1; i <= 20000; i++ {
		larger_input = append(larger_input, i)
	}
	assert.Equal(t, largestRectangleArea([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}), 110)
	assert.Equal(t, largestRectangleArea(larger_input), 110)

}
