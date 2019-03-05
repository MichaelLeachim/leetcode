// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 14:34 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, calculateTokenize("1 + 1"), []int{1, OP_PLUS, 1})
	assert.Equal(t, calculateTokenize(" 2-1 + 2 "), []int{2, OP_MINUS, 1, OP_PLUS, 2})
	assert.Equal(t, calculateTokenize("(1+(4+5+2)-3)+(6+8)"), []int{
		OP_LEFT, 1, OP_PLUS, OP_LEFT, 4, OP_PLUS, 5, OP_PLUS,
		2, OP_RIGHT, OP_MINUS, 3, OP_RIGHT, OP_PLUS, OP_LEFT,
		6, OP_PLUS, 8, OP_RIGHT})

	assert.Equal(t, calculateTokenize(" 21239-111 + 2987 "), []int{21239, OP_MINUS, 111, OP_PLUS, 2987})

	assert.Equal(t, 2, calculate("1 + 1"))
	assert.Equal(t, 3, calculate(" 2-1 + 2 "))
	assert.Equal(t, -1, calculate(" 2-(1 + 2) "))
	assert.Equal(t, -1, calculate(" 2-((((((1 + 2)))))) "))
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
	assert.Equal(t, 0, calculate("0"))
}
