// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-26 15:41 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving: https://leetcode.com/problems/minimum-window-substring/
// This problem does not look hard, just:
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

// When solution? when every letter in pattern appears at least once.

func minWindow(s string, t string) string {
	LEN_OF_PATTERN := len(t)
	LEN_OF_INPUT := len(s)

	curbest := []rune{}
	curiter := []rune{}
	metitem := map[rune]int{}
	bestlen := LEN_OF_INPUT
	itemsappeared := 0
	// initialize store
	for _, rune := range t {
		metitem[rune] = 0
	}
	nowright := true
	leftp := 0
	rightp := 0
	for rightp < LEN_OF_INPUT-1 {
		rsr := rune(s[rightp])
		rsl := rune(s[leftp])
		if nowright {
			if rsl == rsl && leftp != rightp {
				nowright = !nowright
			}
			// advance right pointer
			metcount, ok := metitem[rsr]
			// letter is not a part of pattern
			if !ok {
				curiter = append(curiter, rsr)
				rightp += 1
				continue
			} else {
				// letter is a part of pattern
				if metcount == 0 {
					itemsappeared += 1
				}
				metitem[rsr] += 1
			}
		} else {
			// advance left pointer
			metcount, ok := metitem[rsl]
			// letter is not a part of pattern
			if !ok {
				curiter = curiter[1:]
				leftp += 1
				continue
			} else {
				// letter is a part of pattern
				metcount -= 1
				metitem[rsl] = metcount

				// when touches zero element
				if metcount == 0 {
					itemsappeared -= 1
					nowright = !nowright
				}

			}
		}

	}

}
