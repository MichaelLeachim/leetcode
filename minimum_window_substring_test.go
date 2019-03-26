// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-26 15:46 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindowFailingCase(t *testing.T) {
	assert.Equal(t, "A", minWindow("A", "A"))

}

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "CBA", minWindow("ADOBECBA", "ABC"))

	assert.Equal(t, "C", minWindow("CACACACA", "C"))
	assert.Equal(t, "", minWindow("BBBBBBBBBBBBBBBBBB", "AAA"))
	assert.Equal(t, "A", minWindow("ABBABBAABBBBB", "A"))

	assert.Equal(t, "AA", minWindow("AA", "AA"))
	assert.Equal(t, "aaa", minWindow("aaa", "aaa"))
	assert.Equal(t, "a", minWindow("aaa", "a"))
	assert.Equal(t, "aa", minWindow("aaa", "aa"))

	assert.Equal(t, "aba", minWindow("abbbaaababa", "aba"))

	assert.Equal(t, "A", minWindow("A", "A"))
	assert.Equal(t, "", minWindow("", "A"))
	assert.Equal(t, "", minWindow("", ""))

}
