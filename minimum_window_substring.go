// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-26 15:41 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/minimum-window-substring/
// Have two pointers. Have one go through the S, if it sees EOL, break, if it sees symbol in T increment its counter, check
// whether second pointer stays on this symbol. If it does, go forward decrementing each symbol in T counter, until it equals 0, then, continue from A
// When decrementing or incrementing, if itemsappeared >= len(t) (means, all symbols are in place), check len of curiter. If it is smaller than the best, set best to curiter
// items appeared, decrement when metitem[item] becomes 0, increment, when it becomes 1 from 0.

// That is a solution, but there is a lot of bookeeping. Let's see if we can generalize it.
// While no solution, move right pointer to the right.
// While solution, move left pointer to the left
// ADOBECBA
// ADOBEC BA  (has solution) (no left moving, no solution)
// ADOBECB A  =//=
// ADOBECBA   (left pointer)
// ADOBE CBA  (has solution), best solution CBA

// zaaba, aa
// When exactly solution? when every letter in pattern appears exactly once.
// When no solution? when there is a letter that appears 0 times, but, a letter can be duplicated, so:
//   When no solution? when a letter appears less times than it appears in pattern

// Okay, now, it turns out, the pattern can have non unique letters.
// (Btw, They could have provided an example IN BOLD mentioning it)

// How quickly check if we have a solution?
// If the pattern is unique, the solution would be the appearance of each letter in pattern at least once.
// babzaa aab
// (b|2 a|1) (a|2 b|1)
// We cannot see the difference by sum: because sum_1==sum_2
// But, we can mainain a set of 1|0 matches, and its length equals length of unique items in input
// When do we flip the switch? when item length equal amount of this item on input

// efficiently answers to the question whether current input contains pattern
// whether it will contain that pattern when increment a letter
// whether it will contain that pattern when decrement a letter
type minimumWindowSubstringIterator struct {
	_datum    map[rune]int
	_counter  int
	_lendatum int
}

func newMinimumWindowSubstringIterator(pattern string) *minimumWindowSubstringIterator {
	result := map[rune]int{}
	for _, v := range pattern {
		result[v] += 1
	}
	return &minimumWindowSubstringIterator{_datum: result, _counter: 0, _lendatum: len(result)}
}

func (m *minimumWindowSubstringIterator) dec(a rune) bool {

	_, ok := m._datum[a]
	if !ok {
		return m.stat()
	}

	if m._datum[a] == 0 {
		m._counter -= 1
	}
	m._datum[a] += 1
	return m.stat()

}
func (m *minimumWindowSubstringIterator) stat() bool {
	return m._lendatum == m._counter
}

func (m *minimumWindowSubstringIterator) incString(a string) bool {
	for _, v := range a {
		m.inc(v)
	}
	return m.stat()
}
func (m *minimumWindowSubstringIterator) decString(a string) bool {
	for _, v := range a {
		m.dec(v)
	}
	return m.stat()
}

func (m *minimumWindowSubstringIterator) incStringP(a string) *minimumWindowSubstringIterator {
	m.incString(a)
	return m
}
func (m *minimumWindowSubstringIterator) decStringP(a string) *minimumWindowSubstringIterator {
	m.decString(a)
	return m
}

func (m *minimumWindowSubstringIterator) inc(a rune) bool {
	_, ok := m._datum[a]
	if !ok {
		return m.stat()
	}
	m._datum[a] -= 1
	if m._datum[a] == 0 {
		m._counter += 1
	}
	return m.stat()
}

// While no solution, move right pointer to the right.
// While solution, move left pointer to the left
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	leftp := 0
	rightp := 0
	best := ""
	b := newMinimumWindowSubstringIterator(t)
	for {
		if b.stat() {
			// set up best solution
			if len(best) == 0 || rightp-leftp < len(best) {
				best = s[leftp:rightp]
			}
			if len(best) == len(t) {
				return best
			}
			b.dec(rune(s[leftp]))
			leftp += 1
			continue
		}
		if rightp >= len(s) {
			break
		}
		b.inc(rune(s[rightp]))
		rightp += 1
	}
	return best
}
