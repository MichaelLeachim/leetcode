// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:51 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import ()

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

// solution(a)   => [a]
// solution(aa)  => [a a,aa]
// solution(aaaa)  => [a,aa,aaaa]

// solution(aaa) => [solution(aa) + a, solution(a) + aa]
// suffixOf(aaaa) => [a,aa,aaaa]

// Let's recursively define solution function
// solution(x) = solution(x without suffix) + suffix  for every suffix in matchSufix(x)

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
			// add result without solution
			if len(x)-len(suffix) == 0 {
				result = append(result, suffix)
				continue
			}
			for _, sub_solution := range solution(x[:len(x)-len(suffix)]) {
				result = append(result, sub_solution+" "+suffix)
			}

		}
		solutions[x] = result
		return result
	}

	return solution(s)
}
