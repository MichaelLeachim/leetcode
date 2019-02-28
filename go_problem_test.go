// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 20:51 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func somethingFunkyGoesOn(data [][]string) []string {
	resultStrings := []string{}
	for _, item := range data {
		row := []byte{}
		for _, subItem := range item {
			row = append(row, []byte(subItem)...)
			row = append(row, []byte(" ")...)
		}
		resultStrings = append(resultStrings, string(row[:len(row)-1]))
	}
	return resultStrings
}

// this was smth ephemeral
func TestByteProblems(t *testing.T) {
	input := []string{"a a a a a a a", "aa a a a a a", "a aa a a a a", "a a aa a a a", "aa aa a a a", "aaaa a a a", "a a a aa a a", "aa a aa a a", "a aa aa a a", "a aaaa a a", "a a a a aa a", "aa a a aa a", "a aa a aa a", "a a aa aa a", "aa aa aa a", "aaaa aa a", "a a aaaa a", "aa aaaa a", "a a a a a aa", "aa a a a aa", "a aa a a aa", "a a aa a aa", "aa aa a aa", "aaaa a aa", "a a a aa aa", "aa a aa aa", "a aa aa aa", "a aaaa aa", "a a a aaaa", "aa a aaaa", "a aa aaaa"}

	another_input := [][]string{[]string{"a", "a", "a", "a", "a", "aa", "a"}, []string{"aa", "a", "a", "a", "a", "a"}, []string{"a", "aa", "a", "a", "a", "a"}, []string{"a", "a", "aa", "a", "a", "a"}, []string{"aa", "aa", "a", "a", "a"}, []string{"aaaa", "a", "a", "a"}, []string{"a", "a", "a", "aa", "a", "a"}, []string{"aa", "a", "aa", "a", "a"}, []string{"a", "aa", "aa", "a", "a"}, []string{"a", "aaaa", "a", "a"}, []string{"a", "a", "a", "aa", "aa", "a"}, []string{"aa", "a", "a", "aa", "a"}, []string{"a", "aa", "a", "aa", "a"}, []string{"a", "a", "aa", "aa", "a"}, []string{"aa", "aa", "aa", "a"}, []string{"aaaa", "aa", "a"}, []string{"a", "a", "aaaa", "a"}, []string{"aa", "aaaa", "a"}, []string{"a", "a", "a", "a", "a", "aa"}, []string{"aa", "a", "a", "aa", "aa"}, []string{"a", "aa", "a", "aa", "aa"}, []string{"a", "a", "aa", "aa", "aa"}, []string{"aa", "aa", "a", "aa"}, []string{"aaaa", "a", "aa"}, []string{"a", "a", "a", "aa", "aa"}, []string{"aa", "a", "aa", "aa"}, []string{"a", "aa", "aa", "aa"}, []string{"a", "aaaa", "aa"}, []string{"a", "a", "a", "aaaa"}, []string{"aa", "a", "aaaa"}, []string{"a", "aa", "aaaa"}}
	assert.Equal(t, input, somethingFunkyGoesOn(another_input))
}
