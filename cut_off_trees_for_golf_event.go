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
func sliceGet(some_slice []int, some_index int) (int, bool) {
	if some_index < 0 {
		return -1, false
	}
	if some_index > len(some_slice) {
		return -1, false
	}
	return some_slice[some_index], true
}

func adjacencyMatrix(forest [][]int) map[int]map[int]bool {
	matrix := map[int]map[int]bool{}
	// row_len := len(forest[0])
	// col_len := len(forest)

	// for row_index, row := range forest {
	// 	for col_index, col := range row {
	// 		left, ok := sliceGet(forest, row_index-1)
	// 		right := forest[row_index+1]
	// 		top := forest[col_index-1]
	// 		bottom := forest[col_index+1]

	// 	}
	// }

	return matrix
}

func cutOffTree(forest [][]int) int {
	return 0
}
