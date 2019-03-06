// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 15:21 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"log"
)

// Solving: https://leetcode.com/problems/largest-rectangle-in-histogram/ problem

// 1 2 3 4 5

// solution(1) = {1,1}
// solution(2) = {1,2} or {2,2}
// solution(3) = {1,3} or {2,4} or {3,3}
// solution(4) = {1,4} or {2,6} or {3,6} or {4,4}
// solution(5) = {1,5} or {2,8} or {3,9} or {4,8} or {5,5}

// 5 4 3 2 1

// solution(5) = {5,5}
// solution(4) = {5,5} or {4,8}
// solution(3) = {5,5} or {4,8} or {3,9}
// solution(2) = {5,5} or {4,8} or {3,9} or {2,8}
// solution(1) = {5,5} or {4,8} or {3,9} or {2,8} or {1,5}

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
// solution(2) = {1,4} or {2,6} or {3,3}

// solution(1) = {1,5}

// 5 4 3 4 5
// solution(5) = {5,5}
// solution(4) = {5,5} {4,8}
// solution(3) = {5,5} {4,8}, {3,9}
// solution(4) = {5,5} {4,8}, {3,9}, {4,4}
// solution(5) = {5,5} {4,8}, {3,9}, {4,8},{5,5}

func largestRectangleArea(heights []int) int {
	best := 0
	stack := [][2]int{}
	new_stack := [][2]int{}
	var prev_item [2]int
	for _, cur_height := range heights {
		cur_area := cur_height

		for len(stack) > 0 {
			prev_item, stack = stack[len(stack)-1], stack[:len(stack)-1]

			stack_height, stack_area := prev_item[0], prev_item[1]

			if stack_height < cur_height {
				new_stack = append([][2]int{[2]int{stack_height, stack_area + stack_height}}, new_stack...)
				if stack_area+stack_height > best {
					best = stack_area + stack_height
				}
				continue
			}
			new_stack = append(new_stack, [2]int{stack_height, new_stack_area})

		}
		for i := 0; i < len(stack); i++ {
			// for i := len(stack) - 1; i >= 0; i-- {

			log.Println(stack_height)
			if stack_height < cur_height {
				new_stack_area := stack_area + stack_height
				if new_stack_area > best {
					best = new_stack_area
				}
				// new_stack =

				continue
			}

			if stack_height >= cur_height {
				cur_area += cur_height
			}
		}
		new_stack = append(new_stack, [2]int{cur_height, cur_area})
		stack = new_stack
		log.Println(stack)
		if cur_area > best {
			best = cur_area
		}

	}

	return best

}
