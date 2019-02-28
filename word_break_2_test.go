// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 15:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// Let's define matchSuffix function: matchSuffix(input) will return all suffixes for a given input that match dictionary.
// For example, matchSuffix(sand)  => [and, sand]
//              matchSuffix(sando) => []

func TestMatchSuffix(t *testing.T) {
	wordDict := map[string]bool{"cat": true, "cats": true, "and": true, "sand": true, "dog": true}
	assert.Equal(t, matchSuffix("catsand", wordDict), []string{"and", "sand"})
	assert.Equal(t, matchSuffix("catsando", wordDict), []string{})
	assert.Equal(t, matchSuffix("", wordDict), []string{})
	assert.Equal(t, matchSuffix("pine", map[string]bool{"pine": true}), []string{"pine"})
}

func TestWordBreak2(t *testing.T) {

	// first example
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	assert.Equal(t, wordBreak(s, wordDict), []string{
		"cats and dog",
		"cat sand dog",
	})
	// debugging
	assert.Equal(t, wordBreak("catsanddo", wordDict), []string{})
	assert.Equal(t, wordBreak("catsandog", wordDict), []string{})

	// second example
	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	assert.Equal(t, []string{
		"pine apple pen apple",
		"pineapple pen apple",
		"pine applepen apple",
	}, wordBreak(s, wordDict))
	// debugging
	// solution(pineapple) = [pine apple, pineapple]
	//   suffixes: pineapple, apple
	//
	// solution(pine) = [pine]
	s = "pineapple"
	assert.Equal(t, []string{"pine apple", "pineapple"}, wordBreak(s, wordDict))

	// third example
	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	assert.Equal(t, []string{}, wordBreak(s, wordDict))
}

func BenchmarkWordBreak2(t *testing.B) {
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	result := []string{}
	for i := 0; i <= 10; i++ {
		result = append(result, wordDict[i%len(wordDict)])
	}
	largerDataSet := strings.Join(result, "")
	for n := 0; n < t.N; n++ {
		wordBreak(largerDataSet, wordDict)
	}

}
