// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 15:21 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Solving: https://leetcode.com/problems/largest-rectangle-in-histogram/ problem

// 1 2 3 4 5

// solution(1) = {1,1}
// solution(2) = {1,2} or {2,2}
// solution(3) = {1,3} or {2,4} or {3,3}
// solution(4) = {1,4} or {2,6} or {3,6} or {4,4}
// solution(5) = {1,5} or {2,8} or {3,9} or {4,8} or {5,5}

// 5 4 3 2 1

// solution(5) = {5,5}
// solution(4) = {4,8}
// solution(3) = {3,9}
// solution(2) = {2,8}
// solution(1) = {1,5}

// 5 5 5 5 5 5
// solution(5) = {5,5}
// solution(5) = {5,10},
// solution(5) = {5,15},
// solution(5) = {5,20},
// solution(5) = {5,25},

// 1 2 3 2 1

// solution(1) = {1,1}
// solution(2) = {1,2} or {2,2}
// solution(3) = {1,3} or {2,4} or {3,3}
// solution(2) = {1,4} or {2,6}
// solution(1) = {1,5}

// 5 6 2 3
// solution(5) = {5,5}
// solution(6) = {5,10} {6,6}
// solution(2) = {2,6}
// solution(3) = {2,8}, {3,3}

// 1 2 0 2 1

// solution(1) = {1,1}
// solution(2) = {1,2} {2,2}
// solution(0) = {}
// solution(2) = {2,2}
// solution(1) = {1,2}

// 1 2 1 2 1

// solution(1) = {1,1}
// solution(2) = {1,2} {2,2}
// solution(1) = {1,3}
// solution(2) = {1,4} {2,2}
// solution(1) = {1,5}

// 5 4 3 4 5
// solution(5) = {5,5}
// solution(4) = {4,8}
// solution(3) = {3,9}
// solution(4) = {3,12} {4,4}
// solution(5) = {3,15} {4,8}, {5,1}

//                      0 1 2 3 4
// Given X as index in {5 4 3 4 5}, let's define Solution(Xi) => [{height,area}] function
// For example, Solution(3) => [{3,12} {4,4}]
// Solution(0) => {X[0],X[0]*1} => {5,5}
// Solution(1) => if len(Solution(0)) == 0 return [{height_1,area_1}]
//                for prev_solution in Solution(0):
//                  if(prev_solution_height>=cur_solution_height)
//                    cur_area *= ((prev_area/prev_height) +1)
//                  if prev_solution_height<cur_solution_height
//                    stack[i] = {height_i,area_i+height_i}
//                  append(cur_solution,cur_area)

func largestRectangleArea(heights []int) int {
	stack := [][2]int{}

	best := 0

	for index, height := range heights {
		if height > best {
			best = height
		}

		if index == 0 || len(stack) == 0 {
			stack = append(stack, [2]int{height, height})
			continue
		}
		if height == 0 {
			stack = [][2]int{}
			continue
		}

		new_stack := [][2]int{}
		highest_area_multiplier := 1

		for _, item := range stack {
			item_height, item_area := item[0], item[1]
			if height <= item_height {
				mult := ((item_area / item_height) + 1)
				if mult > highest_area_multiplier {
					highest_area_multiplier = mult
				}
				continue
			}
			// 6/1 + 10/2 + 12/3 + 12/4 + 10/5 + 1
			// 6   + 5    + 4    + 3    + 2    + 1
			// [[1 6] [2 10] [3 12] [4 12] [5 10] [6 6]]

			new_item_area := item_height + item_area
			if new_item_area > best {
				best = new_item_area
			}
			new_stack = append(new_stack, [2]int{item_height, new_item_area})
		}
		stack = new_stack
		cur_area := height * highest_area_multiplier
		if cur_area > best {
			best = cur_area
		}
		stack = append(stack, [2]int{height, cur_area})
	}

	return best

}
