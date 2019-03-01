// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-01 17:27 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringRepetition(t *testing.T) {
	iterator := newStringRepetition("acb", 4)
	result := ""
	for next, hasNext := iterator.next(); hasNext; {
		result += string(next)
	}
	assert.Equal(t, result, "acbacbacbacb")

}

func TestCalculateRepetitions(t *testing.T) {

	assert.Equal(t, calculateRepetitions("acbacbacbacb", "aba"), 2)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "a"), 4)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "b"), 4)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "ba"), 3)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", ""), -1)

	assert.Equal(t, calculateRepetitions("acbacbacbacb", "acbacbacbacb"), 1)

	assert.Equal(t, calculateRepetitions("acbacbacbacb", "acb"), 4)

	// edge cases
	assert.Equal(t, calculateRepetitions("", "bab"), -1)
	assert.Equal(t, calculateRepetitions("", ""), -1)
	assert.Equal(t, calculateRepetitions("bab", ""), -1)

}
