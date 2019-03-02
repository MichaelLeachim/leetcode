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
	pattern := []rune("acb")
	repeater := 4
	lenPattern := len(pattern)
	stringSize := lenPattern * repeater
	result := ""
	for i := 0; i < stringSize; i++ {
		result += string(pattern[i%lenPattern])
	}
	assert.Equal(t, result, "acbacbacbacb")
}

func TestGetMaxRepetitions(t *testing.T) {

	assert.Equal(t, getMaxRepetitions("acb", 4, "aba", 1), 2)
	assert.Equal(t, getMaxRepetitions("acb", 4, "a", 1), 4)
	assert.Equal(t, getMaxRepetitions("acb", 4, "b", 1), 4)
	assert.Equal(t, getMaxRepetitions("acb", 4, "ba", 1), 3)
	assert.Equal(t, getMaxRepetitions("acb", 4, "", 1), -1)

	assert.Equal(t, getMaxRepetitions("acb", 2, "acb", 4), 0)
	assert.Equal(t, getMaxRepetitions("acb", 4, "acb", 4), 1)

	assert.Equal(t, getMaxRepetitions("acb", 4, "acb", 1), 4)

	// edge cases
	assert.Equal(t, getMaxRepetitions("", 1, "bab", 1), -1)
	assert.Equal(t, getMaxRepetitions("", 1, "", 1), -1)
	assert.Equal(t, getMaxRepetitions("bab", 1, "", 1), -1)

}
