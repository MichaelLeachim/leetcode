// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-10 16:49 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSkyline(t *testing.T) {
	assert.Equal(t,
		[][]int{
			[]int{2, 10},
			[]int{3, 15},
			[]int{7, 12},
			[]int{12, 0},
			[]int{15, 10},
			[]int{20, 8},
			[]int{24, 0},
		}, getSkyline([][]int{
			[]int{2, 9, 10},
			[]int{3, 7, 15},
			[]int{5, 12, 12},
			[]int{15, 20, 10},
			[]int{19, 24, 8}}))

	assert.Equal(t, getSkyline([][]int{
		[]int{2, 9, 10},
		[]int{9, 12, 10},
		[]int{12, 15, 10}}), [][]int{[]int{2, 10}, []int{15, 0}})
	assert.Equal(t, getSkyline([][]int{
		[]int{1, 2, 1},
		[]int{1, 2, 2},
		[]int{1, 2, 3},
	}), [][]int{[]int{1, 3}, []int{2, 0}})

}

func TestIterateBeginEndOverSkyLine(t *testing.T) {
	pairs := [][]int{}

	iterateBeginEndOverSkyLine([][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8}},
		func(is_ending bool, building []int) {
			s := "STARTS "
			if is_ending {
				s = "ENDS "
			}
			t.Log(s, building)
			pairs = append(pairs, building)
		})

	assert.Equal(t, pairs, [][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{3, 7, 15},
		[]int{2, 9, 10},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	})

	storedEnding := &MinHeap{}
	heap.Init(storedEnding)
	storedEnding.PushT(1, 5)
	storedEnding.PushT(4, 4)
	storedEnding.PushT(3, 6)
	storedEnding.PushT(2, 2)

	x := []int{}
	for storedEnding.Len() > 0 {
		_, a := storedEnding.PopT()
		x = append(x, a)
	}
	assert.Equal(t, x, []int{2, 4, 5, 6})

	storedEnding2 := &MaxHeap{}
	heap.Init(storedEnding2)
	storedEnding2.PushT(1, 5)
	storedEnding2.PushT(4, 4)
	storedEnding2.PushT(3, 6)
	storedEnding2.PushT(2, 2)
	a, b := storedEnding2.PeekT()
	assert.Equal(t, a, 3)
	assert.Equal(t, b, 6)

	x = []int{}

	for storedEnding2.Len() > 0 {
		_, a := storedEnding2.PopT()
		x = append(x, a)
	}
	assert.Equal(t, x, []int{6, 5, 4, 2})

	storedEnding2.PushT(1, 5)
	storedEnding2.PushT(4, 4)
	storedEnding2.PushT(3, 6)
	storedEnding2.PushT(2, 2)
	storedEnding2.RemoveByRankT(6)
	storedEnding2.RemoveByRankT(4)
	storedEnding2.RemoveByRankT(123)
	x = []int{}
	for storedEnding2.Len() > 0 {
		_, a := storedEnding2.PopT()
		t.Log(a)
		x = append(x, a)
	}
	assert.Equal(t, x, []int{5, 2})

}
