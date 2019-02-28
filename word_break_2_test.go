// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 15:52 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchSuffix(t *testing.T) {
	wordDict := map[string]bool{"cat": true, "cats": true, "and": true, "sand": true, "dog": true}
	assert.Equal(t, matchSuffix("catsand", wordDict), []string{"and", "sand"})
	assert.Equal(t, matchSuffix("catsando", wordDict), []string{})
	assert.Equal(t, matchSuffix("", wordDict), []string{})
}
func TestWordBreak2(t *testing.T) {

	// first example
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	assert.Equal(t, wordBreak(s, wordDict), []string{
		"cats and dog",
		"cat sand dog",
	})

	// second example
	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	assert.Equal(t, wordBreak(s, wordDict), []string{
		"pine apple pen apple",
		"pineapple pen apple",
		"pine applepen apple",
	})

	// third example
	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	assert.Equal(t, wordDict, []string{})

}
