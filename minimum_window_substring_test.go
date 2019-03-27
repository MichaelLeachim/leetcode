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

func TestMinWindowPatternRecognition(t *testing.T) {
	assert.Equal(t, true, minWindowPatternRecognition("babzaa", "aab"))

	assert.Equal(t, true, minWindowPatternRecognition("bbbaaaa", "ab"))
	assert.Equal(t, false, minWindowPatternRecognition("ab", "aaabbbb"))

	assert.Equal(t, true, minWindowPatternRecognition("ba", "ab"))
	assert.Equal(t, true, minWindowPatternRecognition("baa", "aab"))
	assert.Equal(t, true, minWindowPatternRecognition("baaaaaaaaaaaaaaaaaaaa", "aab"))
	assert.Equal(t, false, minWindowPatternRecognition("baaaaaaaaaaaaaaaaaaaaz", "aabzz"))

}

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "A", minWindow("A", "A"))
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "CBA", minWindow("ADOBECBA", "ABC"))

	assert.Equal(t, "C", minWindow("CACACACA", "C"))
	assert.Equal(t, "", minWindow("BBBBBBBBBBBBBBBBBB", "AAA"))
	assert.Equal(t, "A", minWindow("ABBABBAABBBBB", "A"))

	assert.Equal(t, "AA", minWindow("AA", "AA"))

	assert.Equal(t, "aaa", minWindow("aaa", "aaa"))

	assert.Equal(t, "a", minWindow("aaa", "a"))
	assert.Equal(t, "aa", minWindow("aaa", "aa"))
	assert.Equal(t, "aa", minWindow("zaaba", "aa"))
	assert.Equal(t, "aa", minWindow("zaa", "aa"))

	assert.Equal(t, "aabb", minWindow("aabbb", "aabb"))

	assert.Equal(t, "aabb", minWindow("zaabbb", "aabb"))

	assert.Equal(t, "aba", minWindow("abbbaaababa", "aba"))

	assert.Equal(t, "", minWindow("abb", "aba"))

	assert.Equal(t, "A", minWindow("A", "A"))
	assert.Equal(t, "", minWindow("", "A"))
	assert.Equal(t, "", minWindow("", ""))
	assert.Equal(t, "aa", minWindow("zaa", "aa"))
	assert.Equal(t, "aa", minWindow("aaz", "aa"))
	assert.Equal(t, "aa", minWindow("zaaz", "aa"))
	assert.Equal(t, "aa", minWindow("zaazaa", "aa"))
	assert.Equal(t, "aa", minWindow("zaazaa", "aa"))
	assert.Equal(t, "aa", minWindow("aazaa", "aa"))
	assert.Equal(t, "abza", minWindow("babzaa", "aab"))

}
