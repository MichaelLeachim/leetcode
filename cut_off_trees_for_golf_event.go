// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-01 11:23 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

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

// will either return element under this index, and ok == true
//          or return -1                        and ok == false

func canSlice(list_len int, some_index int) bool {
	if some_index < 0 || some_index >= list_len {
		return false
	}
	return true
}

// representation of adjacent graph through adjacency matrix representation
func adjacencyMatrix(forest [][]int) map[int]map[int]bool {

	matrix := map[int]map[int]bool{}
	row_len := len(forest[0])
	col_len := len(forest)

	for row_index, row := range forest {
		for col_index, col := range row {
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

func cutOffTree(forest [][]int) int {
	return 0
}
