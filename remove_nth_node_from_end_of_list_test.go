// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-04 19:30 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	var mklist func(...int) *ListNode
	var list2slice func(*ListNode) []int
	mklist = func(data ...int) *ListNode {
		if len(data) == 0 {
			return nil
		}
		return &ListNode{Val: data[0], Next: mklist(data[1:]...)}
	}
	list2slice = func(item *ListNode) []int {
		if item == nil {
			return []int{}
		}
		return append([]int{item.Val}, list2slice(item.Next)...)
	}
	assert.Equal(t, list2slice(mklist(1, 2, 3, 4, 5)), []int{1, 2, 3, 4, 5})
	assert.Equal(t, list2slice(mklist(1)), []int{1})
	assert.Equal(t, list2slice(mklist()), []int{})

	assert.Equal(t, list2slice(removeNthFromEnd(mklist(1, 2, 3, 4, 5), 2)), []int{1, 2, 3, 5})
	assert.Equal(t, list2slice(removeNthFromEnd(mklist(1, 2, 3, 4, 5), 1)), []int{1, 2, 3, 4})
	assert.Equal(t, list2slice(removeNthFromEnd(mklist(1, 2, 3, 4, 5), 5)), []int{2, 3, 4, 5})

	assert.Equal(t, list2slice(removeNthFromEnd(mklist(1, 2, 3, 4, 5), 0)), []int{1, 2, 3, 4, 5})

	// assert.Equal(t, list2slice(removeNthFromEnd(mklist(1, 2, 3, 4, 5), 2)), []int{1, 2, 3, 5})

}
