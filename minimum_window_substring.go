// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-26 15:41 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import ()

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

// aaa aa
//

// When solution? when every letter in pattern appears at least once.

func minWindow(s string, t string) string {
	LEN_OF_PATTERN, LEN_OF_INPUT := len(t), len(s)
	itemsappeared := 0
	metitem := map[rune]int{}
	mustcount := map[rune]int{}
	// initialize store
	for _, rune := range t {
		metitem[rune] = 0
		mustcount[rune] += 1
	}
	leftp, rightp, bestpatternlen := 0, 0, LEN_OF_INPUT+1
	bestpattern, curpattern := []rune{}, []rune{}

	for {
		if leftp >= LEN_OF_INPUT {
			break
		}
		rsl := rune(s[leftp])
		// if has solution, advance left
		if itemsappeared == LEN_OF_PATTERN {
			lcurpat := len(curpattern)

			if lcurpat == LEN_OF_PATTERN {
				return string(curpattern)
			}
			// bookkeep best pattern
			if bestpatternlen > lcurpat {
				bestpattern = curpattern
				bestpatternlen = lcurpat
			}
			curpattern = curpattern[1:]
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
			if c == 0 {
				itemsappeared += 1
			}
			metitem[rsr] += 1
		}
		rightp += 1
	}

	return string(bestpattern)

}
