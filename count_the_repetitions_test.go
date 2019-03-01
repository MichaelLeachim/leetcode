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

func TestCalculateRepetitions(t *testing.T) {

	assert.Equal(t, calculateRepetitions("acbacbacbacb", "aba"), 0)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "a"), 0)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "b"), 0)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", ""), 0)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "acbacbacbacb"), 0)
	assert.Equal(t, calculateRepetitions("acbacbacbacb", "acb"), 4)

	// edge cases
	assert.Equal(t, calculateRepetitions("", "bab"), 0)
	assert.Equal(t, calculateRepetitions("", ""), 0)
	assert.Equal(t, calculateRepetitions("bab", ""), 0)

}
