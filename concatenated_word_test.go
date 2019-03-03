// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-02 20:42 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestConstructTrie(t *testing.T) {
	trie := constructTrie([]string{"hello", "hella", "helli"})
	assert.Equal(t, len(trie), 8)

	for _, item := range trie {
		fom := ""
		for letter, letter_val := range item {
			if letter_val == -1 {
				fom += " _ "
				continue
			}
			fom += " " + string(ascii_table[letter]) + " "
		}
		log.Println(fom)
	}

	assert.Equal(t, []string{"helli", "hello", "hella"}, queryTrie(0, trie))
	assert.Equal(t, []string{"elli", "ello", "ella"}, queryTrie(1, trie))
	assert.Equal(t, []string{"i", "o", "a"}, queryTrie(4, trie))

	assert.Equal(t, []string{}, queryTrie(12, constructTrie([]string{})))
	assert.Equal(t, []string{}, queryTrie(10, constructTrie([]string{"a"})))
	assert.Equal(t, []string{"a"}, queryTrie(0, constructTrie([]string{"a"})))
	assert.Equal(t, []string{"a", "b"}, queryTrie(0, constructTrie([]string{"a", "a", "a", "a", "b"})))

}
