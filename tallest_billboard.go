// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-28 07:07 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"log"
	// "sort"
)

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

// how to optimize?
// We don't need(?) to store both sides.
// solution(1,2,3,4,5) =
//   {1,2,3,5,6,4,7,9,10,5,9...}

/// 1,2,3
//  solution(X) = for every solution(X-1) as sol:
//                  add X to it. if it contains within prev solutions, it is in best, otherwize, add it
//                if X within prev sums, it is in best

func tallestBillboard(rods []int) int {
	solution := map[int]bool{}
	best := 0
	// sort.Ints(rods)
	for i := 0; i < len(rods); i++ {
		rod := rods[i]

		newsolslice := []int{}
		for sol, _ := range solution {
			newsol := sol + rod
			metsol := solution[newsol]
			if newsol > best && metsol {
				best = newsol
			}
			if !metsol {
				newsolslice = append(newsolslice, newsol)
			}
		}
		for _, newsol := range newsolslice {
			solution[newsol] = true
		}
		solution[rod] = true
		log.Println(rod, solution)

	}
	return best

}
