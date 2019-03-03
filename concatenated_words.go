// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Problem: https://leetcode.com/problems/concatenated-words/

// Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

// Compound words are always longer than simple words that they are part of
// sorted(x)
// Solution(x) returns ([simple words of x], [compound words of x])
// Solution(x) = Solution(x-1) + if(x in simple words) => do nothing
//                                                     => check against simple words,
//                                                          if match, then to [concatenaded words]
//                                                          otherwise      to [simple words]
//  Tricky usecase for greedy check: ["catsdog","cat","sdog","dog"], if match "dog", then the "catsdog" is not going to be matched
// check("catdog") => 1
// check("catdo") =>  0
// check("catd")  =>  0
// check("cat")   =>  1

// check(word) =>
// if empty word:
//   return true
// for suffix in trieMatchAllSuffixes(word):
//   if(suffix == word):
//     return true
//   return check(word - suffix)

// [catsdogscat,cats,cat,dog]
// getAllPrefixes(catsdogscat) => [cats,catsdogscat]
// getAllWordsThatStartWith(cats) => [cats,catsdogscat]

// (cats, catsdogscat)

// checkFn(X) =>
//   for wordsThatStartWithX in getAllWordsThatStartWith(X):
//     if (len(wordThatStartWithX) == len(X)) || checkFn(wordThatStartWithX without X):
//       return true
//   return false

// words that start with X: for each
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

func findAllConcatenatedWordsInADict(words []string) []string {
	// store words in an array
	wordStore := map[string]bool{}
	for _, word := range words {
		wordStore[word] = true
	}

	checkFnMemo := map[string]bool{}
	var checkFn func(string) bool
	checkFn = func(word string) bool {
		checked, ok := checkFnMemo[word]
		if ok {
			return checked
		}
		for _, prefix := range wordsThatStartWith(word, wordStore) {
			if len(prefix) == len(word) || checkFn(word[len(prefix):]) {
				checkFnMemo[word] = true
				return true
			}
		}
		checkFnMemo[word] = false
		return false
	}

	result := []string{}
	for _, word := range words {
		if checkFn(word) {
			result = append(result, word)
		}
	}
	return result
}
