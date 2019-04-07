// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-04-07 19:31 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateDifferenceMatrix(t *testing.T) {
	N := 0x7FF0000000000000
	assert.Equal(t,
		[][]int{
			[]int{0, 0, 2, -3, -3, 0, -2, 1},
			[]int{N, 0, 2, -3, -3, 0, -2, 1},
			[]int{N, N, 0, -5, -5, -2, -4, -1},
			[]int{N, N, N, 0, 0, 3, 1, 4},
			[]int{N, N, N, N, 0, 3, 1, 4},
			[]int{N, N, N, N, N, 0, -2, 1},
			[]int{N, N, N, N, N, N, 0, 3},
			[]int{N, N, N, N, N, N, N, 0}},
		best_time_to_buy_and_sell_stock.GenerateDifferenceMatrix([]int{3, 3, 5, 0, 0, 3, 1, 4}))

	assert.Equal(t, [][]int{
		[]int{0, 2, 1, 3},
		[]int{N, 0, -1, 1},
		[]int{N, N, 0, 2},
		[]int{N, N, N, 0},
	}, best_time_to_buy_and_sell_stock.GenerateDifferenceMatrix([]int{1, 3, 2, 4}))
	assert.Equal(t, [][]int{[]int{0}}, best_time_to_buy_and_sell_stock.GenerateDifferenceMatrix([]int{1}))
	assert.Equal(t, [][]int{}, best_time_to_buy_and_sell_stock.GenerateDifferenceMatrix([]int{}))
}

func TestMaxProfitMatrix(t *testing.T) {
	N := 0x7FF0000000000000

	assert.Equal(t,
		[][]int{
			[]int{0, 0, 2, 2, 2, 2, 2, 2},
			[]int{N, 0, 2, 2, 2, 2, 2, 2},
			[]int{N, N, 2, 2, 2, 2, 2, 2},
			[]int{N, N, N, 2, 2, 3, 3, 4},
			[]int{N, N, N, N, 2, 3, 3, 4},
			[]int{N, N, N, N, N, 3, 3, 4},
			[]int{N, N, N, N, N, N, 3, 4},
			[]int{N, N, N, N, N, N, N, 4},
		}, best_time_to_buy_and_sell_stock.MaxProfitMatrix([]int{3, 3, 5, 0, 0, 3, 1, 4}))

}

func TestMaxProfitValue(t *testing.T) {
	assert.Equal(t, 6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
	assert.Equal(t, 0, maxProfit([]int{}))
	assert.Equal(t, 0, maxProfit([]int{1}))
}
