// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-31 08:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

// Problem definition:
// for a1...aN choose four numbers a,b,c,d, where  index(a)<index(b)<index(c)<index(d) in such a way
//   that (b-a) + (c-d) is maximum

// Can we reuse solution for X-1?
//  [1,3,1,3] => 4 (3-1) + (3-1)
//  solution([1]) => 0   (hardcoded)
//  solution([1,3]) => 2 (hardcoded)
//  solution([1,3,1]) => 2 (hardcoded)
//  solution([1,3,1,3]) => 4
//  solution([1,3,1,3,5]) => 6
//  solution([(1),(3),(1),3,(5),1]) => 6
//  solution([(1),3,1,3,(5),(1),(7)]) => 10

// solution([1,3,2,4]) => 4
// solution matrix
// [0, 2,  1,  3
//  N, 0, -1,  1
//  N, N,  0,  2
// 	N, N,  N,  0 ]

// 2,3
//   1
// 4
// [0, 2,  2,  3
//  N, 0,  2,  2
//  N, N,  2,  2
// 	N, N,  N,  2 ]

// solution([3,3,5,0,0,3,1,4])
// [0  0  2 -3 -3  0 -2  1]
// [N  0  2 -3 -3  0 -2  1]
// [N  N  0 -5 -5 -2 -4 -1]
// [N  N  N  0  0  3  1  4]
// [N  N  N  N  0  3  1  4]
// [N  N  N  N  N  0 -2  1]
// [N  N  N  N  N  N  0  3]
// [N  N  N  N  N  N  N  0]

// So, solution_matrix[bought_day][sold_day] = difference
// As per the exampple:
//  Buy on day 4 (price = 0) and sell on day 6 (price = 3),  solution_matrix[4-1][6-1] = 3
//  Then buy on day 7 (price = 1) and sell on day 8 (price = 4), solution_matrix[7-1][8-1] = 3

// What is the solution?
//  solution_matrix[x][y] + solution_matrix[x+x1][y+y1] = maximum
//  solution_matrix[x-x1][y-y1] + solution_matrix[x][y] = maximum where x1,y1 > 0

// So, maxmatrix[x][y] = max(maxmatrix[x-1],maxmatrix[x-1][y-1],maxmatrix[y-1])

// [0  0  2  2  2  2  2  2]
// [N  0  2  2  2  2  2  2]
// [N  N  2  2  2  2  2  2]
// [N  N  N  2  2  3  3  4]
// [N  N  N  N  2  3  3  4]
// [N  N  N  N  N  3  3  4]
// [N  N  N  N  N  N  3  4]
// [N  N  N  N  N  N  N  4]

type best_time_to_buy_and_sell_stock_ struct {
}

var best_time_to_buy_and_sell_stock = best_time_to_buy_and_sell_stock_{}

func (b *best_time_to_buy_and_sell_stock_) GenerateDifferenceMatrix(prices []int) [][]int {
	fs := len(prices)
	result := make([][]int, fs)
	for i := 0; i < fs; i++ {
		pointdiff := make([]int, fs)
		ps := prices[i]
		for j := 0; j < fs; j++ {
			if j < i {
				pointdiff[j] = 0x7FF0000000000000
			} else {
				pointdiff[j] = prices[j] - ps
			}
		}
		result[i] = pointdiff
	}
	return result
}
func (b *best_time_to_buy_and_sell_stock_) MaxProfitMatrix(prices []int) [][]int {
	bo, _ := b.MaxProfit(prices)
	return bo
}
func (b *best_time_to_buy_and_sell_stock_) MaxProfitValue(prices []int) int {
	_, bo := b.MaxProfit(prices)
	return bo
}

// naive implementation
func (b *best_time_to_buy_and_sell_stock_) MaxProfit(prices []int) ([][]int, int) {
	maxmatrix := b.GenerateDifferenceMatrix(prices)
	fs := len(prices)
	best := 0
	for i := 0; i < fs; i++ {
		for j := i; j < fs; j++ {
			right := maxmatrix[i][j]
			lessbest := 0
			for i1 := 0; i1 < i; i1++ {
				for j1 := j; j1 < j; j1++ {
					if lessbest < maxmatrix[i1][j1] {
						lessbest = maxmatrix[i1][j1]
					}
				}
			}
			if right+lessbest > best {
				best = right + lessbest
			}
		}
	}
	return maxmatrix, best

}

func (b *best_time_to_buy_and_sell_stock_) MaxProfitLater(prices []int) ([][]int, int) {
	maxmatrix := b.GenerateDifferenceMatrix(prices)
	best := 0

	fs := len(prices)
	for i := 0; i < fs; i++ {
		for j := i; j < fs; j++ {
			comp := 0
			cur := maxmatrix[i][j]
			// if i > 0 {
			// 	comp = maxmatrix[i-1][j]
			// }

			// if j > i {
			// 	s := maxmatrix[i][j-1]
			// 	if comp < s {
			// 		comp = s
			// 	}
			// }

			if i > 0 && j > i {
				s := maxmatrix[i-1][j-1]
				if comp < s {
					comp = s
				}
				if cur+s > best {
					best = cur + s
				}
			}
			if comp > cur {
				maxmatrix[i][j] = comp
			}
		}
	}
	return maxmatrix, best
}

func maxProfit(prices []int) int {
	return best_time_to_buy_and_sell_stock.MaxProfitValue(prices)
}
