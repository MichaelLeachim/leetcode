// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 15:21 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Solving: https://leetcode.com/problems/largest-rectangle-in-histogram/ problem

// solution(X,h) =>
//   if(X < 0): return 0
//   if(h>data[x]): return 0
//   return solution(X-1,h) + h

// Okay, that is, although working, is not an intended solution
// Should think more on this problem

// What is the largest rectangle?
//   It is a rectagle, whose height * width is maximum
//   For given width, largest rectangle is that, whose height is maximum
///  For given heigh, largest rectangle is that, which contains maximum width

// best for every index
// 5 8 9 8 5
// 1 2 3 4 5

func largestRectangleArea(heights []int) int {
	var largestRectangleAreaSolution func(heights []int, x int, h int) int
	combineDigits := func(x, y uint32) uint64 {
		return uint64(x)<<32 | uint64(y)
	}
	memoized := map[uint64]int{}
	largestRectangleAreaSolution = func(heights []int, x int, h int) int {
		if x < 0 || h > heights[x] {
			return 0
		}

		if x == 0 {
			return h
		}

		cmb := combineDigits(uint32(x), uint32(h))
		s, ok := memoized[cmb]
		if ok {
			return s
		}
		memoized[cmb] = largestRectangleAreaSolution(heights, x-1, h) + h
		return memoized[cmb]
	}

	len_height := len(heights)
	if (len_height) == 0 {
		return 0
	}
	best := 0
	seen_heights := map[int]bool{}
	for index, height := range heights {
		// add simle heuristics ti minimize search?
		seen_heights[height] = true

		// no solutions before, so do not count
		if index-1 < 0 {
			best = height
			continue
		}
		// if prev height is smaller than current, no point in calculating
		// upper bounds of current height
		prev_height := heights[index-1]
		if height > prev_height && height > best {
			best = height
		}

		// no point in calculating heights that haven't been seen before (???)
		for i := range seen_heights {
			if i > height {
				continue
			}
			s := largestRectangleAreaSolution(heights, index, i)
			if s > best {
				best = s
				break
			}
		}

	}
	return best
}

func largestRectangleAreaSolution(heights []int, x int, h int) int {
	if x < 0 || h > heights[x] {
		return 0
	}
	if x == 0 {
		return h
	}
	return largestRectangleAreaSolution(heights, x-1, h) + h
}
