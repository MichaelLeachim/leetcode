// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:51 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Working on this problem:
// https://leetcode.com/problems/word-break-ii/

// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]

// Let's define a function `solution` that, given the input,
// returns the amount of existing solutions.

// solution(catsanddog) = 2
// solution(catsanddo) = 2
// solution(catsandd ) = 2
// solution(catsand  ) = 2 (cat sand, cats and)
// solution(catsan   ) = 0
// solution(catsa    ) = 0
// solution(cats     ) = 2
// solution(cat      ) = 1
// solution(ca       ) = 0
// solution(c        ) = 0
// solution(         ) = 0

// So, let's recursively define the solution function
// solution(x) = if match(suffix_0) != 0 then  match(suffix_0) + solution(len(x)-len(suffix_0)),
//                match(suffix_1) + solution(len(x)-len(suffix_1))
//                match(suffix_N) + solution(len(x)-len(suffix_N)))

// where match(suffix) = len(wordDict[suffix])
