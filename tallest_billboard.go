// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-28 07:07 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
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

// solution(X) cannot be larger  (than sum of all up to X)/2

// Several implementation details:
//  * appending to a slice is a costly operation, we should initialize the slice of required size beforehand
//  * storing data is a pain point, I am used to recursive implementation, iterative is hard for me

//

// How to make iterative implementation easier?
//  have two slices, for left part and for the right part
// Okay, how to make recursive implementation?
//   * Write the solution(X) function:
//     for i in solution(X-1):
//       append to result [left_i+rod,right_i]
//       append to result [left_i,    right_i+rod]
//     append to result
//     return result

// Solution(X) = max(Solution(X-1), Solution(X-1) + right, Solution(X-1) + left)
// Okay, what solutions are guaranteed to not to converge?
//  * Solution, where left||right part is larger than the (sum of all)/2
//  * Solution, where best must be

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
	sumarray := []int{rods[0]}
	for i := 1; i < len(rods); i++ {
		sumarray = append(sumarray, sumarray[i-1]+rods[i])
	}
	// maxsum := 0
	// for _, rod := range rods {
	// 	maxsum += rod
	// }
	// maxsum /= 2

	best := 0
	for rodindex, rod := range rods {
		lensol := len(solution)
		for i := 0; i < lensol; i++ {
			left, right := solution[i][0], solution[i][1]

			rr, lr := right+rod, left+rod
			if rr == left && left > best {
				best = left
			}
			if lr == right && right > best {
				best = right
			}
			if rr <= sumarray[rodindex] {
				nda([2]int{left, rr})
			}
			if lr <= sumarray[rodindex] {
				nda([2]int{lr, right})
			}
		}
		nda([2]int{0, rod})
		nda([2]int{rod, 0})
	}
	return best

}
