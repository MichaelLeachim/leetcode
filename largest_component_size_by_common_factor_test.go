// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-08 18:57 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestComponentSize(t *testing.T) {
	assert.Equal(t, largestComponentSize([]int{2, 3, 6}), 3)
	assert.Equal(t, largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}), 8)
	assert.Equal(t, largestComponentSize([]int{20, 50, 9, 63}), 2)
	assert.Equal(t, largestComponentSize([]int{4, 6, 15, 35}), 4)
	// assert.Equal(t, largestComponentSize([]int{}), 0)

	assert.Equal(t, largestComponentSize([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}), 0)
	assert.Equal(t, largestComponentSize([]int{2, 2, 2, 2, 2, 2, 2}), 7)
	assert.Equal(t, largestComponentSize([]int{4, 6, 8, 10, 12, 14, 15}), 7)

	assert.Equal(t, largestComponentSize([]int{3, 5, 7, 11, 13, 17, 19}), 1)

}
