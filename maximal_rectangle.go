// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-19 11:18 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving https://leetcode.com/problems/maximal-rectangle/ problem

// Solution for one param of dynamic programming:
// For row X, get all possible rectangles for row X-1
// For every rectangle in row X-1, get all matches in X: for example
// 1 1 1  (X-1)
// 0 1 1  (X)
// So, the resulting match is:
// 1 1
// 1 1
// Plus, add all rectangles in X, as if there is no X-1. For example
// 0 0 0 0
// 1 1 0 1
// Will add rectangles: 1 1 and 1

// Dynamic programming solution for two parameters:
// Solution for row X (and col Y) is solution for X-1 (and col Y) + solution in X (and col Y)
//   It is enough to find all solutions only in the corner matrix[X][Y] because we will be calculating results for every point in matrix
//    if corner == "0": no solutions
//    currentSolution = solution(Y-1) with height = 1 + "1"
//    get all X-1 solutions
//    for every solution, if it is smaller than currentSolution, increment it by its width
//    for every solution, that is larger than currentSolution, multiply its height by X
//    for every solution, that is equal currentSolution, increment it by currentSolution and set (currentSolution) to nil
//    return currentSolution + mapped X-1 solutions
// Can we reuse Y-1 solutions?

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	// ONE := []byte("1")[0]
	ZERO := []byte("0")[0]
	MATRIX_HEIGHT := len(matrix)
	MATRIX_WIDTH := len(matrix[0])

	// initialize storage
	memo := make([][][][2]int, MATRIX_HEIGHT)
	levelsol := make([][]int, MATRIX_HEIGHT)
	for i := 0; i < MATRIX_HEIGHT; i++ {
		memo[i] = make([][][2]int, MATRIX_WIDTH)
		levelsol[i] = make([]int, MATRIX_WIDTH)
	}
	maxarea := 0

	for rowindex, row := range matrix {
		for colindex, col := range row {
			if col == ZERO {
				memo[rowindex][colindex] = [][2]int{}
				continue
			}

			cursolwidth := 1
			if colindex-1 >= 0 {
				cursolwidth = levelsol[rowindex][colindex-1] + 1
			}
			cursolheight := 1
			cursols := [][2]int{}

			if rowindex-1 >= 0 {
				for _, solution := range memo[rowindex-1][colindex] {
					solwidth, solheight := solution[0], solution[1]
					if solwidth <= cursolwidth {
						solarea := solwidth * (solheight + 1)
						if solarea > maxarea {
							maxarea = solarea
						}
						cursols = append(cursols, [2]int{solwidth, solheight + 1})
						continue
					}
					solarea := cursolwidth * (solheight + 1)
					if solarea > maxarea {
						maxarea = solarea
					}
					cursols = append(cursols, [2]int{cursolwidth, solheight + 1})
					continue

				}
			}
			carea := cursolwidth * cursolheight
			if carea > maxarea {
				maxarea = carea
			}

			cursols = append(cursols, [2]int{cursolwidth, cursolheight})

			levelsol[rowindex][colindex] = cursolwidth
			memo[rowindex][colindex] = cursols
		}
	}

	return maxarea
}
