// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-26 15:41 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"log"
)

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

// But, if pattern is not unique?
// aab, aabb a  a2|b1 = 2
// Well, we can do similar thing:
//   calculate unique letters in pattern [a,b]
//   calculate amount each letter appears in pattern [a|2,b|1]
//   decrement counter by one, when the amount is lower, than the amount of this letter in a pattern
//   increment counter, by one, when amount is higher than the amount of this letter in a pattern

// aab, aabb
// solution(aab) = solution(aab)[a|2 b|1] + b

// Naive solution:
//  for every letter in curinput
//    for every match, increment match by one.
//    if match equals amount in pattern, increment sum
//    if match more than amount in pattern, do nothing
//
//    if it is in the pattern, decrement in pattern by one
//
// Recursive solution: Solution(baa) = Solution(ba) + a = {b:1,a:1} + a = {b:1,a:2}

func minWindowPatternRecognition(s string, t string) bool {
	dat := map[rune]int{}
	for _, v := range t {
		dat[v] += 1
	}
	reso := 0
	for _, v := range s {
		_, ok := dat[v]
		if ok {
			dat[v] -= 1
			if dat[v] == 0 {
				reso += 1
			}

		}
	}
	log.Println(dat, reso)
	return reso == len(dat)
}

func minWindow(s string, t string) string {
	_, LEN_OF_INPUT := len(t), len(s)
	itemsappeared := 0
	metitem := map[rune]int{}
	mustcount := map[rune]int{}
	// initialize store
	for _, rune := range t {
		metitem[rune] = 0
		mustcount[rune] += 1
	}
	LEN_OF_PATTERN_UNIQUE := len(mustcount)

	leftp, rightp, bestpatternlen := 0, 0, LEN_OF_INPUT+1
	bestpattern, curpattern := []rune{}, []rune{}

	for leftp < LEN_OF_INPUT {
		rsl := rune(s[leftp])
		// if has solution, advance left
		if itemsappeared == LEN_OF_PATTERN_UNIQUE {
			log.Println(string(curpattern), t, itemsappeared)
			lcurpat := len(curpattern)

			// if lcurpat == LEN_OF_PATTERN {
			// 	return string(curpattern)
			// }
			// bookkeep best pattern
			if bestpatternlen > lcurpat {
				bestpattern = curpattern
				bestpatternlen = lcurpat
			}
			if len(curpattern) != 0 {
				curpattern = curpattern[1:]
			}

			leftp += 1
			c, ok := metitem[rsl]
			if ok {
				if c == mustcount[rsl] {
					itemsappeared -= 1
				}
				metitem[rsl] -= 1
			}
			continue
		}

		if rightp >= LEN_OF_INPUT {
			break
		}
		rsr := rune(s[rightp])

		// if has no solution, advance right
		curpattern = append(curpattern, rsr)

		// bookkeep solution
		c, ok := metitem[rsr]
		if ok {
			if c == mustcount[rsl] {
				itemsappeared += 1
			}
			metitem[rsr] += 1
		}
		rightp += 1
	}

	return string(bestpattern)

}
