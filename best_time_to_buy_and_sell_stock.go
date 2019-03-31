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

// So, now our task is to choose two numbers from the solution matrix
// in such a way, that index(number A)<index(number B), column index(A)<column index(B) and (number A + number B) = max
// meaning, we can either walk down, or right.
// when walking right, choose maximum
// when walking down,  either choose maximum from the left,
// TODO, figure out this algo
// TODO2: set next problem to be edit distance: https://leetcode.com/problems/edit-distance/

// [(1-1), (3-1),  (2-1),  (4-1)
//  N,     (3-3),  (2-3),  (4-3)
//  N,      N,     (2-2)   (4-2)
// 	N,      N,      N,     (4-4) ]
