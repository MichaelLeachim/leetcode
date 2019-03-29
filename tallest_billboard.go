// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-28 07:07 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import ()

// Solving: https://leetcode.com/problems/tallest-billboard/
// [1,2,3,6] => {1,2,3} {6}
// Find two largest sums that are equal (cannot reuse members)

// solution(1) = {1} {}, {} {1}
// solution(2) = solution(1) + 2 on left, solution(1) + 2 on right + empty solution
// solution(2) = {1,2} {}, {2} {1},  {1} {2}, {} {1,2}  {2}{}, {}{2}
// solution(3) = solution(2) + 3 on left, solution(2) + 3 on right
// solution(3) = {1,2,3} {}, {2,3} {1},  {1,3} {2}, {3,3} {1,2}  {2,3}{}, {3}{2}
//               {1,2} {3}, {2} {1,3},  {1} {2,3}, {} {1,2,3}  {2}{3}, {}{2,3} {3}{}, {},{3}

// It seems, that we must maintain 1,2,4,8,16,32,64,128,256, 2^N variants, where N is 1000, is unpractical
// oh, my mistake, the max N is 20, thus, the problem is solvable, but we must maintain 1048576 variants
// That is quite a lot

// Can we prune some of the solutions?
//  if left or right part is >2500, then there is no solution (by definition)
//  TODO: think something more on it

// Several implementation details:
//  * appending to a slice is a costly operation, we should initialize the slice of required size beforehand
//  * storing data is a pain point, I am used to recursive implementation, iterative is hard for me

// How to make iterative implementation easier?
//  have two slices, for left part and for the right part
// Okay, how to make recursive implementation?
//   * Write the solution(X) function:
//     for i in solution(X-1):
//       append to result [left_i+rod,right_i]
//       append to result [left_i,    right_i+rod]
//     append to result
//     return result

// How to make iterative implementation?

// solution(2, 1, 2) =>
// 1,2,2
// {2}{0} {0}{2}
// {2,1}{0} {0,1}{2} {2}{0,1} {0}{2,1} {1}{0} {0}{1}
// {2,1,2}{0} {1,2}{2} {2,2}{1} {2}{2,1} {1,2}{} {2}{1}
// {2,1}{2} {1}{2,2} {2}{1,2} {}{2,1,2} {1}{2} {}{1,2} {2}{} {}{2}

func tallestBillboard(rods []int) int {
	solution := [][2]int{}
	nodup := map[[2]int]bool{}

	// performant append

	nda := func(sol [2]int) {
		if !nodup[sol] {
			solution = append(solution, sol)
			nodup[sol] = true
		}
	}

	best := 0
	for _, rod := range rods {
		for _, item := range solution {
			left, right := item[0], item[1]
			rr, lr := right+rod, left+rod
			if rr == left && left > best {
				best = left
			}
			if lr == right && right > best {
				best = right
			}
			if rr <= 2500 {
				nda([2]int{left, rr})
			}
			if lr <= 2500 {
				nda([2]int{lr, right})
			}
		}
		nda([2]int{0, rod})
		nda([2]int{rod, 0})
	}
	return best

}
