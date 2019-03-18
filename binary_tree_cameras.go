// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-13 12:11 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/binary-tree-cameras/

// For root node, check amount of colored children if the root is colored and if the root is uncolored. Choose minimum
// Node **must be colored** when not parent nor any child is colored
// Can we simplify? Node must be colored when parent is not colored? No
// Another approach:
// * We either color given node or not, unless we must.
// * We must color the node, when parent node is uncolored and children nodes are not colored.
// * We choose to color or not to, based on score (amount of colored nodes) down the tree

// Solve(node, parentColored?) =>
//  if (node == nil) {
//    if parentColored? return 0,false
//    else              return 1,true
//  }
//  when_self_colored,child_colored? := Solve(node.Left,true) + Solve(node.Right,true)
//  when_self_not_colored,any_child_colored? := Solve(node.fLeft,false) + Solve(node.Right,false)
//  reso := min(when_self_colored+1,when_self_not_colored)
//  must color item, cannot avoid
//  if not any_child_colored? and not parentColored?:
//    return when_self_colored+1
//  return amount,true

type binary_tree_cameras_ struct {
	_memoized_if_true  map[*TreeNode][2]int
	_memoized_if_false map[*TreeNode][2]int
}

var binary_tree_cameras = binary_tree_cameras_{
	_memoized_if_true:  map[*TreeNode][2]int{},
	_memoized_if_false: map[*TreeNode][2]int{},
}

func (b *binary_tree_cameras_) memoize(node *TreeNode, parentColored bool, count int, isColored bool) {
	isColoredInt := -1
	if isColored {
		isColoredInt = 1
	}

	if parentColored {
		b._memoized_if_true[node] = [2]int{isColoredInt, count}
		return
	}
	b._memoized_if_false[node] = [2]int{isColoredInt, count}
}

func (b *binary_tree_cameras_) unmemoize(node *TreeNode, parentColored bool) (bool, int, bool) {
	var res [2]int
	var ok bool

	if parentColored {
		res, ok = b._memoized_if_true[node]
	} else {
		res, ok = b._memoized_if_false[node]
	}

	if !ok {
		return false, -1, false
	}

	is_colored_int, count := res[0], res[1]
	is_colored := false
	if is_colored_int == 1 {
		is_colored = true
	}
	return true, count, is_colored

}

func (b *binary_tree_cameras_) tts(t *TreeNode) []*TreeNode {
	if t == nil {
		return []*TreeNode{}
	}
	return append(append([]*TreeNode{t}, b.tts(t.Left)...), b.tts(t.Right)...)
}

func (b *binary_tree_cameras_) SolveIterative(t *TreeNode) int {
	nodes := b.tts(t)
	solution := make([]int, len(nodes))
	binsol := make([]bool, len(nodes))

	for pos, t := range nodes {
		if t == nil {
			solution[pos] = 0
			binsol[pos] = false
			continue
		}

		// if leaf
		if t.Left == nil && t.Right == nil {
			if binsol[pos-1] {
				solution[pos] = 0
				binsol[pos] = false
				continue
			}
			solution[pos] = 1
			binsol[pos] = true
			continue
		}
	}

	self_colored_node_left, _ := b.Solve(t.Left, true)
	self_colored_node_rigt, _ := b.Solve(t.Right, true)
	self_colored_node_score := self_colored_node_left + self_colored_node_rigt

	self_uncolored_node_left, is_child_colored_left := b.Solve(t.Left, false)
	self_uncolored_node_rigt, is_child_colored_right := b.Solve(t.Right, false)
	self_uncolored_node_score := self_uncolored_node_left + self_uncolored_node_rigt

	// node should be colored by the score
	if self_colored_node_score < self_uncolored_node_score {
		b.memoize(t, parentColored, self_colored_node_score+1, true)
		return self_colored_node_score + 1, true
	}

	// !(false || true || true)
	// node should be colored by the rules
	if (parentColored || is_child_colored_left || is_child_colored_right) == false {
		b.memoize(t, parentColored, self_colored_node_score+1, true)
		return self_colored_node_score + 1, true
	}
	b.memoize(t, parentColored, self_uncolored_node_score, false)
	return self_uncolored_node_score, false

}

func (b *binary_tree_cameras_) Solve(t *TreeNode, parentColored bool) (int, bool) {
	ok, result, result2 := b.unmemoize(t, parentColored)
	if ok {
		return result, result2
	}

	if t == nil {
		return 0, false
	}
	// if leaf
	if t.Left == nil && t.Right == nil {
		if parentColored {
			return 0, false
		}
		return 1, true
	}

	self_colored_node_left, _ := b.Solve(t.Left, true)
	self_colored_node_rigt, _ := b.Solve(t.Right, true)
	self_colored_node_score := self_colored_node_left + self_colored_node_rigt

	self_uncolored_node_left, is_child_colored_left := b.Solve(t.Left, false)
	self_uncolored_node_rigt, is_child_colored_right := b.Solve(t.Right, false)
	self_uncolored_node_score := self_uncolored_node_left + self_uncolored_node_rigt

	// node should be colored by the score
	if self_colored_node_score < self_uncolored_node_score {
		b.memoize(t, parentColored, self_colored_node_score+1, true)
		return self_colored_node_score + 1, true
	}

	// !(false || true || true)
	// node should be colored by the rules
	if (parentColored || is_child_colored_left || is_child_colored_right) == false {
		b.memoize(t, parentColored, self_colored_node_score+1, true)
		return self_colored_node_score + 1, true
	}
	b.memoize(t, parentColored, self_uncolored_node_score, false)
	return self_uncolored_node_score, false

}

func (b *binary_tree_cameras_) CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + b.CountNodes(root.Left) + b.CountNodes(root.Right)
}

func minCameraCover(root *TreeNode) int {

	// binary_tree_cameras._memoized = map[*TreeNode]bool{}
	res, _ := binary_tree_cameras.Solve(root, false)
	return res
}
