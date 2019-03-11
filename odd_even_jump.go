// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-11 16:24 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// solving: https://leetcode.com/problems/odd-even-jump/ problem

// Okay, rules here are not easy.
// Let's translate them into human readable formate

// We can jump **only forward**
// Rules for odd jump:
//  * new item is larger or equal  current item
//  * new item is smaller than any other item excluding current (and all previously visited)
//  * if new item appears multiple times, jump to the rightmost (smallest index) one
// Rules for even jump:
//  * new item is smaller or equal current item
//  * new item is larger than any other item excluding current item (and all previously visited)
//  * if new item appears multiple times, jump to the righmost (smallest index) one

// Let's define a `step` function for odd jump:
//   next_jump := -1
//   next_jump_score := MaxInt
//   previous_smallest := {empty set}
//   for every point in array:
//     if point >= cur_item && point not in visited && point < smallest_point && point not in previous_smallest {
//        next_jump = (index of point)
//        next_jump_score = point
//        previous_smallest[point] = (index of point)
//     }

// And `step` function for even jump:
//   next_jump := -1
//   next_jump_score := 0
//   previous_largest := {empty set}
//   for every point in array:
//     if point <= cur_item && point not in visited && point > smallest_point && point not in previous_largest {
//        next_jump = (index of point)
//        next_jump_score = point
//        previous_smallest[point] = (index of point)
//     }

// Let's define a `solve` function
// solve(odd?, index) =>
//  1. if odd? or even? find best next jump (if none, return false)
//  2. return solve(!odd?,best_next_jump)

func oddEvenJumps(A []int) int {
	// Rules for odd jump:
	//  * new item is larger or equal  current item
	//  * new item is smaller than any other item excluding current
	//  * if new item appears multiple times, jump to the rightmost (smallest index) one
	odd_step := func(index int) int {
		cur_item := A[index]
		next_jump := -1
		next_jump_score := 1<<31 - 1
		previous_smallest := map[int]bool{}
		for point_index := index; point_index < len(A); point_index++ {
			point := A[point_index]
			if point >= cur_item && point < next_jump_score && !previous_smallest[point] {
				next_jump = point_index
				next_jump_score = point
				previous_smallest[point] = true
			}
		}
		return next_jump
	}
	// Rules for even jump:
	//  * new item is smaller or equal current item
	//  * new item is larger than any other item excluding current item (and all previously visited)
	//  * if new item appears multiple times, jump to the righmost (smallest index) one
	even_step := func(index int) int {
		cur_item := A[index]
		next_jump := -1
		next_jump_score := 0
		previous_largest := map[int]bool{}
		for point_index := index; point_index < len(A); point_index++ {
			point := A[point_index]
			if point <= cur_item && point > next_jump_score && !previous_largest[point] {
				next_jump = point_index
				next_jump_score = point
				previous_largest[point] = true
			}
		}
		return next_jump
	}

	return 0

}
