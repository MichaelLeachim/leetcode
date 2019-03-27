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

func TestMinWindowIncrementDecrementSolution(t *testing.T) {
	b := newMinimumWindowSubstringIterator

	// test incrementing
	assert.Equal(t, false, b("aab").stat())
	assert.Equal(t, false, b("aab").incString("a"))
	assert.Equal(t, false, b("aab").incString("aa"))
	assert.Equal(t, true, b("aab").incString("aab"))

	assert.Equal(t, true, b("aab").incString("babzaa"))
	assert.Equal(t, true, b("ab").incString("bbbaaaa"))
	assert.Equal(t, true, b("ab").incString("aaabbbb"))
	assert.Equal(t, true, b("aa").incString("zaaz"))
	assert.Equal(t, true, b("").incString("zaaz"))

	assert.Equal(t, true, b("ba").incString("ab"))
	assert.Equal(t, true, b("baa").incString("aab"))
	assert.Equal(t, true, b("aab").incString("baaaaaaaaaaaaaaaaaaaa"))
	assert.Equal(t, false, b("baaaaaaaaaaaaaaaaaaaa").incString("aab"))

	assert.Equal(t, false, b("baaaaaaaaaaaaaaaaaaaaz").incString("aabzz"))
	assert.Equal(t, true, b("aa").incString("aazaa"))

	// test decrementing
	assert.Equal(t, true, b("abc").incString("adobecodebanc"))
	assert.Equal(t, true, b("abc").incStringP("adobecodebanc").decString("adobecode"))
	assert.Equal(t, false, b("abc").incStringP("adobecodebanc").decString("adobecodeb"))
	assert.Equal(t, true, b("abc").incStringP("adobecodebanc").decString("odebanc"))
	assert.Equal(t, false, b("abc").incStringP("adobecodebanc").decString("codebanc"))

	x := b("aab")
	x.incString("babzaa")
	assert.Equal(t, x.stat(), true)
	assert.Equal(t, true, x.decString("a"))
	assert.Equal(t, false, x.decString("a"))
	assert.Equal(t, true, x.incString("a"))
	assert.Equal(t, true, x.decString("b"))
	assert.Equal(t, false, x.decString("b"))

	assert.Equal(t, true, b("aab").incString("baa"))
	assert.Equal(t, b("aab").incStringP("baa")._datum, map[int32]int{97: 0, 98: 0})
	assert.Equal(t, false, b("aab").incStringP("baa").decString("a"))

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

	assert.Equal(t, "baa", minWindow("abbbaaababa", "aba"))

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

func BenchmarkMinWindow(t *testing.B) {
	for n := 0; n < t.N; n++ {
		minWindow("ADOBECODEBANC", "ABC")
	}
}
