// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-11 16:24 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// solving: https://leetcode.com/problems/odd-even-jump/ problem

// Okay, rules here are not easy.
// Let's translate them into human readable format

// We can jump **only forward**
// Rules for odd jump:
//  * new item is larger or equal  current item
//  * new item is smaller than any other item excluding current (and all previously visited)
//  * if new item appears multiple times, jump to the rightmost (smallest index) one
// Rules for even jump:
//  * new item is smaller or equal current item
//  * new item is larger than any other item excluding current item (and all previously visited)
//  * if new item appears multiple times, jump to the righmost (smallest index) one

// Or, in other words:
// 	*  For odd jumps, get smallest item that is larger or equal current
//  *  For even jumps, get largest item, that is smaller or equal current

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
type odd_even_jump_ struct{}

var odd_even_jump = odd_even_jump_{}

// Rules for odd jump:
//  * new item is larger or equal  current item
//  * new item is smaller than any other item excluding current
//  * if new item appears multiple times, jump to the rightmost (smallest index) one
// For odd jumps, get smallest item that is larger or equal current
func (o *odd_even_jump_) OddStep(index int, A []int) int {
	// in case it is the end
	if len(A)-1 == index {
		return index
	}

	cur_item := A[index]
	best_index := -1
	best_score := 1<<31 - 1
	previous_smallest := map[int]bool{}
	for point_index := index + 1; point_index < len(A); point_index++ {
		point := A[point_index]
		if point >= cur_item && point < best_score && !previous_smallest[point] {
			best_index = point_index
			best_score = point
			previous_smallest[point] = true
		}
	}
	return best_index
}

// Rules for even jump:
//  * new item is smaller or equal current item
//  * new item is larger than any other item excluding current item
//  * if new item appears multiple times, jump to the leftmost (smallest index) one

// For even jumps, get largest item, that is smaller or equal current
func (o *odd_even_jump_) EvenStep(index int, A []int) int {
	// in case it is the end
	if len(A)-1 == index {
		return index
	}
	cur_item := A[index]

	best_index := -1
	best_score := 0
	previous_largest := map[int]bool{}
	for point_index := index + 1; point_index < len(A); point_index++ {
		point := A[point_index]
		if point <= cur_item && point > best_score && !previous_largest[point] {
			best_index = point_index
			best_score = point
			previous_largest[point] = true
		}
	}
	return best_index
}

// solve(odd?, index) =>
//  1. if odd? or even? find best next jump (if none, return false)
//  2. return solve(!odd?,best_next_jump)

func (o *odd_even_jump_) Solve(memoized map[bool]map[int]bool, isOdd bool, item_index int, A []int) bool {
	if item_index == len(A)-1 {
		return true
	}

	if item_index == -1 {
		return false
	}
	res, ok := memoized[isOdd][item_index]
	if ok {
		return res
	}
	var st int
	if isOdd {
		st = o.OddStep(item_index, A)
	} else {
		st = o.EvenStep(item_index, A)
	}
	res = o.Solve(memoized, !isOdd, st, A)
	memoized[isOdd][item_index] = res
	return res

}

func (o *odd_even_jump_) OddEvenJumps(A []int) int {
	memoized := map[bool]map[int]bool{
		true:  map[int]bool{},
		false: map[int]bool{},
	}
	total := 0

	for index, _ := range A {
		if o.Solve(memoized, true, index, A) {
			total += 1
		}
	}
	return total
}

func oddEvenJumps(A []int) int {
	return odd_even_jump.OddEvenJumps(A)

}
