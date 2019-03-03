// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-02 20:42 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordsThatStartWith(t *testing.T) {
	assert.Equal(t, []string{"helloworld", "hello"}, wordsThatStartWith("helloworld", map[string]bool{"hello": true, "helloworld": true, "brabus": true, "holo": true}))
	assert.Equal(t, []string{}, wordsThatStartWith("", map[string]bool{"hello": true, "helloworld": true, "brabus": true, "holo": true}))
	assert.Equal(t, []string{}, wordsThatStartWith("h", map[string]bool{"hello": true, "helloworld": true, "brabus": true, "holo": true}))
	assert.Equal(t, []string{"hello", "helloworld"}, wordsThatStartWith("helloworldblab", map[string]bool{"hello": true, "helloworld": true, "brabus": true, "holo": true}))
	assert.Equal(t, []string{}, wordsThatStartWith("", map[string]bool{}))
	assert.Equal(t, []string{}, wordsThatStartWith("cats", map[string]bool{"cat": true, "cats": true}))
}

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	assert.Equal(t, findAllConcatenatedWordsInADict([]string{"catsdogscat", "cats", "cat", "dog"}), []string{"catsdogscat"})
	assert.Equal(t, findAllConcatenatedWordsInADict([]string{"catdogcat", "cat", "dog"}), []string{"catdogcat"})

	assert.Equal(t, findAllConcatenatedWordsInADict([]string{"catsdog", "cat", "sdog", "dog"}), []string{"catsdog"})
	assert.Equal(t, findAllConcatenatedWordsInADict([]string{}), []string{})
	assert.Equal(t, findAllConcatenatedWordsInADict([]string{"a", "aa", "aaa", "aaaa"}), []string{"aa", "aaaa", "aaa"})
	assert.Equal(t, findAllConcatenatedWordsInADict([]string{"", "aa", "aaa", "aaaa"}), []string{"aaaa"})

}
