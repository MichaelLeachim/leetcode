// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-02 20:42 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestConstructQueryAndTraverseTrie(t *testing.T) {
	// testing construction
	trie := constructTrie([]string{"hello", "hella", "helli"})
	assert.Equal(t, len(trie), 8)
	// visual representation of trie matrix
	for _, item := range trie {
		fom := ""
		for letter, letter_val := range item {
			if letter_val == -1 {
				fom += " _ "
				continue
			}
			fom += " " + string(ascii_table[letter]) + " "
		}
		t.Log(fom)
	}
	t.Log("=========================================================")
	for _, item := range constructTrie([]string{"i", "o", "bbbb"}) {
		fom := ""
		for letter, letter_val := range item {
			if letter_val == -1 {
				fom += " ___ "
				continue
			}
			fom += " " + string(ascii_table[letter]) + "|" + strconv.Itoa(letter_val) + " "
		}
		t.Log(fom)
	}

	// test querying trie
	assert.Equal(t, []string{"helli", "hello", "hella"}, queryTrie(0, trie))
	assert.Equal(t, []string{"elli", "ello", "ella"}, queryTrie(1, trie))
	assert.Equal(t, []string{"i", "o", "a"}, queryTrie(4, trie))
	assert.Equal(t, []string{"i", "o", "bbbb"}, queryTrie(0, constructTrie([]string{"i", "o", "bbbb"})))
	assert.Equal(t, []string{"i"}, queryTrie(0, constructTrie([]string{"i", "i", "i"})))
	assert.Equal(t, []string{}, queryTrie(0, constructTrie([]string{"", "", ""})))

	// test edge cases
	assert.Equal(t, []string{}, queryTrie(12, constructTrie([]string{})))
	assert.Equal(t, []string{}, queryTrie(10, constructTrie([]string{"a"})))
	assert.Equal(t, []string{"a"}, queryTrie(0, constructTrie([]string{"a"})))
	assert.Equal(t, []string{"a", "b"}, queryTrie(0, constructTrie([]string{"a", "a", "a", "a", "b"})))
	assert.Equal(t, []string{}, queryTrie(128, constructTrie([]string{"a", "a", "a", "a", "b"})))
	assert.Equal(t, []string{}, queryTrie(-100, constructTrie([]string{"a", "a", "a", "a", "b"})))
}
