// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-19 12:18 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximalRectangle10(t *testing.T) {
	assert.Equal(t, 10,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
		}))

}

func TestMaximalRectangle(t *testing.T) {
	assert.Equal(t, 6,
		maximalRectangle([][]byte{
			[]byte("10100"),
			[]byte("10111"),
			[]byte("11111"),
			[]byte("10010"),
		}))

	assert.Equal(t, 3,
		maximalRectangle([][]byte{
			[]byte("10111"),
		}))

	assert.Equal(t, 1,
		maximalRectangle([][]byte{
			[]byte("1"),
		}))

	assert.Equal(t, 1,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("0"),
		}))

	assert.Equal(t, 2,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
		}))
	assert.Equal(t, 0,
		maximalRectangle([][]byte{}))

	assert.Equal(t, 4,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
		}))

	assert.Equal(t, 5,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
		}))

	assert.Equal(t, 10,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
		}))
	assert.Equal(t, 3,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("1"),
			[]byte("0"),
		}))
	assert.Equal(t, 2,
		maximalRectangle([][]byte{
			[]byte("11"),
		}))
	assert.Equal(t, 3,
		maximalRectangle([][]byte{
			[]byte("111"),
		}))

	assert.Equal(t, 2,
		maximalRectangle([][]byte{
			[]byte("1"),
			[]byte("1"),
		}))

	assert.Equal(t, 4,
		maximalRectangle([][]byte{
			[]byte("11"),
			[]byte("11"),
		}))
	assert.Equal(t, 9,
		maximalRectangle([][]byte{
			[]byte("111"),
			[]byte("111"),
			[]byte("111"),
			[]byte("000"),
			[]byte("000"),
		}))

}
