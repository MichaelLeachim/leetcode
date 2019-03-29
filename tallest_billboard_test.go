// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-28 07:08 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTallestBillboardFailingCases(t *testing.T) {

	assert.Equal(t, 11, tallestBillboard([]int{5, 5, 1, 1, 1, 5, 5}))

	assert.Equal(t, 10, tallestBillboard([]int{5, 5, 1, 5, 5}))

	assert.Equal(t, 8, tallestBillboard([]int{1, 2, 3, 5, 5}))

	assert.Equal(t, 5, tallestBillboard([]int{1, 2, 3, 4}))

	assert.Equal(t, 5, tallestBillboard([]int{5, 1, 5}))
	assert.Equal(t, 6, tallestBillboard([]int{5, 1, 5, 1}))
	assert.Equal(t, 6, tallestBillboard([]int{5, 1, 1, 5}))
	assert.Equal(t, 11, tallestBillboard([]int{5, 5, 1, 1, 5, 5}))

	assert.Equal(t, 1, tallestBillboard([]int{1, 1, 1}))
	assert.Equal(t, 2, tallestBillboard([]int{1, 1, 1, 1}))
	assert.Equal(t, 2, tallestBillboard([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, 3, tallestBillboard([]int{1, 1, 1, 1, 1, 1}))
	assert.Equal(t, 3, tallestBillboard([]int{1, 1, 1, 1, 1, 1, 1}))
	assert.Equal(t, 5, tallestBillboard([]int{5, 1, 5}))
	assert.Equal(t, 10, tallestBillboard([]int{5, 5, 1, 5, 5}))
	assert.Equal(t, 4, tallestBillboard([]int{2, 2, 1, 2, 2}))
	assert.Equal(t, 2, tallestBillboard([]int{2, 1, 2}))

}
func TestTallestBillboard(t *testing.T) {
	assert.Equal(t, 0, tallestBillboard([]int{1, 2}))

	assert.Equal(t, 6, tallestBillboard([]int{1, 2, 3, 6}))

	assert.Equal(t, 10, tallestBillboard([]int{1, 2, 3, 4, 5, 6}))

	assert.Equal(t, 3, tallestBillboard([]int{3, 3}))
	assert.Equal(t, 5, tallestBillboard([]int{17, 4, 1, 2, 3}))

	assert.Equal(t, 3, tallestBillboard([]int{3, 3}))
	assert.Equal(t, 9, tallestBillboard([]int{4, 5, 6, 3}))
	assert.Equal(t, 16, tallestBillboard([]int{0, 11, 4, 5, 6, 3, 1, 2}))
	assert.Equal(t, 10, tallestBillboard([]int{2, 3, 4, 5, 6}))

	assert.Equal(t, 10, tallestBillboard([]int{2, 3, 4, 5, 6}))

	assert.Equal(t, 10, tallestBillboard([]int{1, 2, 3, 4, 5, 6}))

	assert.Equal(t, 27, tallestBillboard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	assert.Equal(t, 10, tallestBillboard([]int{1, 2, 3, 4, 5, 5}))

	// 2019/03/29 08:25:41 0 2
	// 2019/03/29 08:25:41 1 8
	// 2019/03/29 08:25:41 2 26
	// 2019/03/29 08:25:41 3 80
	// 2019/03/29 08:25:41 4 242
	// 2019/03/29 08:25:41 5 728
	// 2019/03/29 08:25:41 6 2186
	// 2019/03/29 08:25:41 7 6560
	// 2019/03/29 08:25:41 8 19682
	// 2019/03/29 08:25:41 9 59048
	// 2019/03/29 08:25:41 10 177146
	// 2019/03/29 08:25:41 11 531440
	// 2019/03/29 08:25:41 12 1594322
	// 2019/03/29 08:25:41 13 4782968
	// 2019/03/29 08:25:41 14 14348906
	// 2019/03/29 08:25:43 15 43046720
	// 2019/03/29 08:26:13 16 129140162

	assert.Equal(t, 10, tallestBillboard([]int{10, 10}))
	assert.Equal(t, 10, tallestBillboard([]int{5, 5, 5, 5}))
	assert.Equal(t, 3, tallestBillboard([]int{1, 3, 3}))
	assert.Equal(t, 161, tallestBillboard([]int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}))
	assert.Equal(t, 1275, tallestBillboard([]int{137, 142, 161, 144, 160, 153, 141, 136, 131, 147, 155, 165, 140, 135, 172, 165, 166}))

}

func BenchmarkTallestBillboard(t *testing.B) {
	for n := 0; n < t.N; n++ {
		tallestBillboard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	}
}

func BenchmarkTallestBillboard2(t *testing.B) {
	for n := 0; n < t.N; n++ {
		tallestBillboard([]int{137, 142, 161, 144, 160, 153, 141, 136, 131, 147, 155, 165, 140, 135, 172, 165, 166})
	}
}
