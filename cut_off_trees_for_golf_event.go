// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-01 11:23 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"sort"
)

// Working on https://leetcode.com/problems/cut-off-trees-for-golf-event/

// For all tree heights [1,2,3,...N] as k, calculate shortest
// BFS paths between k_n and k_n+1.
// if no shortest path exists, return -1

// optimisations: Floyd-Warshall algorithm will help calculate
// distances between any two nodes in O(N^3) time

// Representing as a graph, where V is a weight of a node.
// And E is an adjacent connection (either top,left,right or bottom)
// Representing it as adjacency matrix, where matrix[V][V] = <true | false>,
// depending if these two are adjacent

// Construction of an adjacency matrix representation of this graph:
// for each node, if it is not 0,
//   for adjacent_node in  [left,right,top and bottom nodes]
//     if exist and not == 0:
//       matrix[node][adjacent_node] = true

// will check if it is possible to return an element from slice, given its index
// returns false in case of bounding errors
func canSlice(list_len int, some_index int) bool {
	if some_index < 0 || some_index >= list_len {
		return false
	}
	return true
}

// representation of adjacent graph through adjacency matrix representation
// Construction of an adjacency matrix representation of this graph:
// for each node, if it is not 0,
//   for adjacent_node in  [left,right,top and bottom nodes]
//     if exist and not == 0:
//       matrix[node][adjacent_node] = true
func adjacencyMatrix(forest [][]int) map[int]map[int]bool {

	matrix := map[int]map[int]bool{}
	row_len := 0
	col_len := len(forest)

	if canSlice(col_len, 0) {
		row_len = len(forest[0])
	}

	for row_index, row := range forest {
		for col_index, col := range row {
			if col == 0 {
				continue
			}
			adjacents := map[int]bool{}

			// left
			if canSlice(row_len, col_index-1) && row[col_index-1] != 0 {
				adjacents[row[col_index-1]] = true
			}

			// right
			if canSlice(row_len, col_index+1) && row[col_index+1] != 0 {
				adjacents[row[col_index+1]] = true
			}

			// top
			if canSlice(col_len, row_index-1) && forest[row_index-1][col_index] != 0 {
				adjacents[forest[row_index-1][col_index]] = true
			}
			// bottom
			if canSlice(col_len, row_index+1) && forest[row_index+1][col_index] != 0 {
				adjacents[forest[row_index+1][col_index]] = true
			}
			matrix[col] = adjacents
		}
	}
	return matrix
}

func getDistanceByBFS(startNode int, targetNode int, matrix map[int]map[int]bool) (int, bool) {
	queue := []int{}
	visited := map[int]bool{}
	distance := map[int]int{}
	var s int

	if startNode == targetNode {
		return 0, true
	}

	queue = append(queue, startNode)
	for len(queue) > 0 {
		// pop operation
		s, queue = queue[len(queue)-1], queue[:len(queue)-1]
		for adjacent, _ := range matrix[s] {
			if visited[adjacent] {
				continue
			}
			visited[adjacent] = true
			distance[adjacent] = distance[s] + 1
			// push operation
			queue = append([]int{adjacent}, queue...)
			// I am not entirely sure that it is true for all cases
			// should think about this case more
			if adjacent == targetNode {
				dst, ok := distance[adjacent]
				return dst, ok
			}
		}
	}
	dst, ok := distance[targetNode]
	return dst, ok
}

// For all tree heights [1,2,3,...N] as k, calculate shortest
// BFS paths between k_n and k_n+1.
// if no shortest path exists, return -1
// Starting position:
//   calculate          when starting position is the smallest
// * when it is not

func cutOffTree(forest [][]int) int {
	// [0,0] point of traversal
	startingPoint := -1
	if canSlice(len(forest), 0) && canSlice(len(forest[0]), 0) {
		startingPoint = forest[0][0]
	}
	if startingPoint == -1 {
		return -1
	}

	// calculate adjacency matrix
	adjMatrix := adjacencyMatrix(forest)

	// sort heights
	order := []int{}
	for node, _ := range adjMatrix {
		order = append(order, node)
	}
	sort.Ints(order)
	if len(order) == 0 {
		return -1
	}

	distance, ok := getDistanceByBFS(startingPoint, order[0], adjMatrix)
	if !ok {
		return -1
	}
	for i := 1; i < len(order); i++ {
		pointA, pointB := order[i-1], order[i]
		pabDistance, ok := getDistanceByBFS(pointA, pointB, adjMatrix)
		if !ok {
			return -1
		}
		distance += pabDistance
	}

	return distance
}
