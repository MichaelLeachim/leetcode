// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-04 19:08 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Problem statement
// 0  1  2  3  4
// 1->2->3->4->5, and n = 2, node to remove = 5-2 = 3
// removal: set .Next of the node, before the node to remove to the Next of the node to remove
//      or: set .Next

// Solution algorithm:
//  get final node_plus_one (that is head node + N+1)
//  for every node:
//    if(final_node_plus_one is nil){
//      current_node.Next =  current_node.Next.Next
//      return
//    }

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	finalNode := head
	for i := 0; i < n; i++ {
		finalNode = head.Next
	}
	for curNode := head; curNode != nil; curNode = curNode.Next {
		if finalNode.Next == nil {
			curNode.Next = curNode.Next.Next
			return head
		}
		finalNode = finalNode.Next
	}
	return head

}
