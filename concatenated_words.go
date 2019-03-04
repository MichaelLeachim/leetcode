// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Problem: https://leetcode.com/problems/concatenated-words/

// Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
// [Note] Compound words are always longer than simple words that they are part of
// sorted(x)
// Solution(x) returns ([simple words of x], [compound words of x])
// Solution(x) = Solution(x-1) + if(x in simple words) => do nothing
//                                                     => check against simple words,
//                                                          if match, then to [concatenaded words]
//                                                          otherwise      to [simple words]
// No point in doing that, although the idea that compound words are always longer wants to be exploited

// [Tricky usecase] for greedy check: ["catsdog","cat","cats","dog"], if match "dog", then the "catsdog" is not going to be matched

// solution(X,isFirstLevel?)
//   for prefix in X:
//     isFirstLevel:
//       prefix is empty => fatal("cannot happen")
//       prefix == X => continue
//     notFirstLevel:
//       prefix is empty => return true
//       prefix == X => return true (match happened)
//     // if no solution found, keep looking. If solution found, stop looking and return result
//     if solution(X-prefix,false) =>  return true
//   return false
func findAllConcatenatedWordsInADict(words []string) []string {
	// store words in an array
	wordStore := map[string]bool{}
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		wordStore[word] = true
	}

	checkFnMemo := map[string]bool{}
	var checkFn func(string, bool) bool
	checkFn = func(word string, isFirstLevel bool) bool {
		isFirstLevelString := "0"
		if isFirstLevel {
			isFirstLevelString = "1"
		}
		checked, ok := checkFnMemo[word+isFirstLevelString]
		if ok {
			return checked
		}
		for _, prefix := range wordsThatStartWith(word, wordStore) {
			if isFirstLevel && len(prefix) == len(word) {
				continue
			}
			if len(prefix) == 0 {
				checkFnMemo[word+isFirstLevelString] = true
				return true
			}
			if len(prefix) == len(word) {
				checkFnMemo[word+isFirstLevelString] = true
				return true
			}
			if checkFn(word[len(prefix):], false) {
				checkFnMemo[word+isFirstLevelString] = true
				return true
			}
		}
		checkFnMemo[word+isFirstLevelString] = false
		return false
	}

	result := []string{}
	for _, word := range words {
		if checkFn(word, true) {
			result = append(result, word)
		}
	}
	return result
}

// words that start with X
func wordsThatStartWith(input string, words map[string]bool) []string {
	result := []string{}
	_, ok := words[input]
	if ok {
		result = append(result, input)
	}
	for i := 0; i < len(input); i++ {
		_, ok := words[input[:i]]
		if ok {
			result = append(result, input[:i])
		}
	}
	return result
}
