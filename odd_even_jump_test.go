// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-11 20:51 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddEvenSteps(t *testing.T) {
	assert.Equal(t, odd_even_jump.OddStep(0, []int{10, 13, 12, 14, 15}), 2)
	assert.Equal(t, odd_even_jump.EvenStep(2, []int{10, 13, 12, 14, 15}), -1)

	assert.Equal(t, odd_even_jump.OddStep(1, []int{10, 13, 12, 14, 15}), 3)
	assert.Equal(t, odd_even_jump.OddStep(2, []int{10, 13, 12, 14, 15}), 3)
	assert.Equal(t, odd_even_jump.EvenStep(3, []int{10, 13, 12, 14, 15}), -1)
	assert.Equal(t, odd_even_jump.OddStep(3, []int{10, 13, 12, 14, 15}), 4)
	assert.Equal(t, odd_even_jump.OddStep(4, []int{10, 13, 12, 14, 15}), 4)

	assert.Equal(t, odd_even_jump.OddStep(0, []int{2, 3, 1, 1, 4}), 1)
	assert.Equal(t, odd_even_jump.EvenStep(1, []int{2, 3, 1, 1, 4}), 2)
	assert.Equal(t, odd_even_jump.OddStep(2, []int{2, 3, 1, 1, 4}), 3)
	assert.Equal(t, odd_even_jump.EvenStep(3, []int{2, 3, 1, 1, 4}), -1)

	assert.Equal(t, odd_even_jump.OddStep(1, []int{2, 3, 1, 1, 4}), 4)
	assert.Equal(t, odd_even_jump.OddStep(2, []int{2, 3, 1, 1, 4}), 3)
	assert.Equal(t, odd_even_jump.OddStep(3, []int{2, 3, 1, 1, 4}), 4)
	assert.Equal(t, odd_even_jump.OddStep(4, []int{2, 3, 1, 1, 4}), 4)
}

func TestOddEvenJumps(t *testing.T) {
	assert.Equal(t, oddEvenJumps([]int{10, 13, 12, 14, 15}), 2)
	assert.Equal(t, oddEvenJumps([]int{2, 3, 1, 1, 4}), 3)
	assert.Equal(t, oddEvenJumps([]int{5, 1, 3, 4, 2}), 3)
	assert.Equal(t, oddEvenJumps([]int{}), 0)
	m := map[bool]map[int]bool{
		true:  map[int]bool{},
		false: map[int]bool{},
	}
	assert.Equal(t, 6, oddEvenJumps([]int{1, 2, 3, 2, 1, 4, 4, 5}))

	assert.Equal(t, 4, odd_even_jump.OddStep(0, []int{1, 2, 3, 2, 1, 4, 4, 5}))

	assert.Equal(t, 3, odd_even_jump.OddStep(1, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 4, odd_even_jump.EvenStep(3, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 5, odd_even_jump.OddStep(4, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 6, odd_even_jump.EvenStep(5, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 7, odd_even_jump.OddStep(6, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, odd_even_jump.Solve(m, true, 1, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)

	assert.Equal(t, -1, odd_even_jump.EvenStep(4, []int{1, 2, 3, 2, 1, 4, 4, 5}))

	assert.Equal(t, 3, odd_even_jump.OddStep(1, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 4, odd_even_jump.EvenStep(3, []int{1, 2, 3, 2, 1, 4, 4, 5}))

	assert.Equal(t, 5, odd_even_jump.OddStep(2, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 6, odd_even_jump.EvenStep(5, []int{1, 2, 3, 2, 1, 4, 4, 5}))
	assert.Equal(t, 7, odd_even_jump.OddStep(6, []int{1, 2, 3, 2, 1, 4, 4, 5}))

	assert.Equal(t, odd_even_jump.Solve(m, true, 0, []int{1, 2, 3, 2, 1, 4, 4, 5}), false)
	assert.Equal(t, odd_even_jump.Solve(m, true, 1, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
	assert.Equal(t, odd_even_jump.Solve(m, true, 2, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
	assert.Equal(t, odd_even_jump.Solve(m, true, 3, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
	assert.Equal(t, odd_even_jump.Solve(m, true, 4, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
	assert.Equal(t, odd_even_jump.Solve(m, true, 5, []int{1, 2, 3, 2, 1, 4, 4, 5}), false)
	assert.Equal(t, odd_even_jump.Solve(m, true, 6, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
	assert.Equal(t, odd_even_jump.Solve(m, true, 7, []int{1, 2, 3, 2, 1, 4, 4, 5}), true)
}

func BenchmarkOddEvenJumpsSmall(t *testing.B) {
	for n := 0; n < t.N; n++ {
		oddEvenJumps([]int{10, 13, 12, 14, 15})
	}
}
