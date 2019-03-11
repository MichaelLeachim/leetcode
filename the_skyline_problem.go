// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-10 13:55 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"container/heap"
)

// Solving: https://leetcode.com/problems/the-skyline-problem/

// I think, the best approach here is to use sweep line technique.

// Sweep line algorithm

// how to maintain ordered beginning/end calculation:
// for every building in left sorted:
//   if heap is empty: push -building[right_point]
//   if building[right_point] is bigger than heap.pop()
//     process(heap.pop())
//   if building[right_point] is smaller than heap.top():
//     push -building[right_point]

//  for every item in beginning/end point:
//    if it is the end and top element of <line_sweep_heap> is smaller or equal <item_height>:
//      remove building from <line_sweep_heap> and set <result_heights> as the .top() of the <line_sweep_heap>
//      if top element of <line_sweep_heap> is larger than <item_height>:
//        find cur element in the heap by its height and remove it
//    if it is the end, and the heap is empty:
//      set result_heights to the 0

//    if it is the beginning, and the heap is empty:
//      set result_heights as item_height
//      push item height to the heap
//    if it is the beginning and .top of heap is larger than item_height:
//      push item_height to the heap
//    if it is the beginning,and the .top of heap is smaller than the item_height
//      set result_heights as item_height
//      push item_height to the heap

// An IntHeap is a min-heap of ints.
// it is a pair of {value,rank}
type MinHeap [][2]int
type MaxHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Len() int           { return len(h) }

func (h *MinHeap) Push(data interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, data.([2]int))
}
func (h *MinHeap) PushT(item, rank int) {
	heap.Push(h, interface{}([2]int{item, rank}))
}
func (h *MinHeap) PopT() (int, int) {
	d := heap.Pop(h).([2]int)
	return d[0], d[1]
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Push(data interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, data.([2]int))
}
func (h *MaxHeap) PushT(item, rank int) {
	heap.Push(h, interface{}([2]int{item, rank}))
}

func (h *MaxHeap) PopT() (int, int) {
	d := heap.Pop(h).([2]int)
	return d[0], d[1]
}

func (h *MaxHeap) PeekT() (int, int) {
	ho := *h
	return ho[0][0], ho[0][1]
}

// remove item by rank
func (h *MaxHeap) RemoveByRankT(rank int) {
	for i, j := range *h {
		if j[1] == rank {
			heap.Remove(h, i)
			return
		}
	}

}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// how to maintain ordered beginning/end calculation:
// for every building in left sorted:
//   if heap is empty: push -building[right_point]
//   if building[right_point] is bigger than heap.pop()
//     process(heap.pop())
//   if building[right_point] is smaller than heap.top():
//     push -building[right_point]

func iterateBeginEndOverSkyLine(buildings [][]int, process func(is_ending bool, building []int)) {
	storedEnding := &MinHeap{}
	heap.Init(storedEnding)
	index := 0

	for index < len(buildings) {
		left, right := buildings[index][0], buildings[index][1]

		// if empty, push right side of the building  onto the heap, process the building
		if storedEnding.Len() == 0 {
			storedEnding.PushT(index, right)
			process(false, buildings[index])
			index += 1
			continue
		}

		// if the rightmost side of the stored item
		// is smaller than the current left side, process it, retaining index
		i, rank := storedEnding.PopT()
		if left > rank {
			process(true, buildings[i])
			continue
		}
		// push back, as it wasn't used
		storedEnding.PushT(i, rank)

		// if the rightmost side of the item on the stack is larger
		// than the current left side, then process the current item with its left side
		// and push its right side on the heap
		process(false, buildings[index])
		storedEnding.PushT(index, right)
		index += 1
	}

	for storedEnding.Len() > 0 {
		i, _ := storedEnding.PopT()
		process(true, buildings[i])
	}
}

//  for every item in beginning/end point:
//    if it is the end and top element of <line_sweep_heap> is smaller or equal <item_height>:
//      remove building from <line_sweep_heap> and set <result_heights> as the .top() of the <line_sweep_heap>
//      if top element of <line_sweep_heap> is larger than <item_height>:
//        find cur element in the heap by its height and remove it
//    if it is the end, and the heap is empty:
//      set result_heights to the 0

//    if it is the beginning, and the heap is empty:
//      set result_heights as item_height
//      push item height to the heap
//    if it is the beginning and .top of heap is larger than item_height:
//      push item_height to the heap
//    if it is the beginning,and the .top of heap is smaller than the item_height
//      set result_heights as item_height
//      push item_height to the heap

func getSkyline(buildings [][]int) [][]int {
	var left int
	var right int
	var height int
	resultHeights := [][]int{}
	// initialize heap
	storedHeight := &MaxHeap{}
	heap.Init(storedHeight)
	append2res := func(left, height int) {
		if len(resultHeights) == 0 {
			resultHeights = append(resultHeights, []int{left, height})
			return
		}
		p := resultHeights[len(resultHeights)-1]
		prevLeft, prevHeight := p[0], p[1]

		if prevHeight == height {
			return
		}
		if prevLeft == left {
			resultHeights = append(resultHeights[:len(resultHeights)-1], []int{left, height})
			return
		}
		resultHeights = append(resultHeights, []int{left, height})
		return
	}
	// iterate over buildings
	iterateBeginEndOverSkyLine(buildings, func(is_ending bool, building []int) {
		left, right, height = building[0], building[1], building[2]
		is_beginning := !is_ending

		if is_beginning {

			storedHeight.PushT(-1, height)

			_, maxHeight := storedHeight.PopT()
			if height == maxHeight {
				append2res(left, maxHeight)
			}
			storedHeight.PushT(-1, maxHeight)
			return
		}

		if is_ending {
			_, maxHeight := storedHeight.PopT()

			// this item height is lower than the current maximum height
			// so, remove it from the heap & push back maxHeight onto the heap
			if maxHeight > height {
				storedHeight.RemoveByRankT(height)
				storedHeight.PushT(-1, maxHeight)
				return
			}

			// if item height is equal to current maximum height (means, we popped the top item)
			// we set the height to the largest height on the heap (or to 0, if the .Len of the heap is 0)
			if storedHeight.Len() == 0 {
				append2res(right, 0)

				return
			}
			_, sts := storedHeight.PeekT()
			append2res(right, sts)
			return
		}
	})
	return resultHeights
}
