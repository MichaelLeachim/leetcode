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

// solution(catsanddog) = [cats and dog, cat sand dog]
// solution(catsanddo) =  []
// solution(catsandd ) =  []
// solution(catsand  ) =  [cat sand, cats and]
// solution(catsan   ) =  []
// solution(catsa    ) =  []
// solution(cats     ) =  [cats]
// solution(cat      ) =  [cat]
// solution(ca       ) =  []
// solution(c        ) =  []
// solution(         ) =  []

// Let's define match function: match(suffix) => wordDict[suffix] or Nothing
// Let's recursively define solution function:
// solution(x) =   [each + " " + suffix for each in solution(x[len(x)-len(suffix):])]  when len(x)-len(suffix) > 0 and match(suffix)
//                 [suffix]                                                                 len(x)-len(suffix) = 0 and match(suffix)
//                 []

// Let's define matchSuffix function: matchSuffix(input) will return all suffixes for a given input that match dictionary.
// For example, matchSuffix(sand)  => [and, sand]
//              matchSuffix(sando) => []

// Let's recursively define solution function
// solution(x) = [sub_solution+ " " + suffix for sub_solution in solution(x without suffix) when sub_solution is not empty]  for every suffix in matchSufix(x)  if len(x) - len(suffix) > 0
//               matchSufix(x)                                                                                                                                  if len(x) - len(suffix) = 0

func matchSuffix(input string, wordDictAsDict map[string]bool) []string {
	result := []string{}
	for i := len(input); i >= 0; i-- {
		_, ok := wordDictAsDict[input[i:]]
		if ok {
			result = append(result, input[i:])
		}
	}
	return result
}

func wordBreak(s string, wordDict []string) []string {
	var solution func(x string) []string

	// prep wordDict for faster access
	wordDictAsDict := map[string]bool{}
	for _, word := range wordDict {
		wordDictAsDict[word] = true
	}
	solutions := map[string][]string{}

	solution = func(x string) []string {
		res, ok := solutions[x]
		if ok {
			return res
		}

		result := []string{}
		for _, suffix := range matchSuffix(x, wordDictAsDict) {
			if (len(x) - len(suffix)) == 0 {
				result = append(result, suffix)
			}
			for _, sub_solution := range solution(x[:len(x)-len(suffix)]) {
				if len(sub_solution) > 0 {
					result = append(result, sub_solution+" "+suffix)
				}
			}
		}
		return result
	}
	return solution(s)
}
