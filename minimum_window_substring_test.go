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

func TestMinWindow(t *testing.T) {
	assert.Equal(t, minWindow("ADOBECODEBANC", "ABC"), "BANC")
	assert.Equal(t, minWindow("ADOBECBA", "ABC"), "CBA")

	assert.Equal(t, minWindow("CACACACA", "C"), "")
	assert.Equal(t, minWindow("BBBBBBBBBBBBBBBBBB", "AAA"), "")
	assert.Equal(t, minWindow("ABBABBAABBBBB", "A"), "")

}
