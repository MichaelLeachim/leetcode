// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:54 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Working on https://leetcode.com/problems/count-the-repetitions/

// Given:
// S1 = acbacbacbacbacb
// S2 = abab

// Find maximum M, such as S2*M is within S1
// Where within, means, it contains the same symbols in the same order.
// For example:
// S1 = acbacbacbacbacb
// S2 = ab
// Matches: (a)c(b) (a)c(b) (a)c(b) (a)c(b) (a)c(b), total: 5

// Let amount of matches be K
// Approaches. Minimum M is 1, the K is the largest possible for this S2
//             Maximum M is N, where there is only 1 K

// Continuing with the given example:
// K is 5
// Minimum M is 1, so the K is 5, and the string is: (ab)
// Maximum M is 5, so the K is 1, and the string is: (ab)(ab)(ab)(ab)(ab)

// Continuing with another example:
// S1 = acbacbacbacbacb
// S2 = abab
// K = 2, Maximum M is 2

// Okay, if I am not mistaken, this problem looks trivial.
// Where are the possible pitfalls?
// I don't see here any Dynamic Programming, that is why I am afraid, I am missing something.

// Okay, let's solve it, and look at the possible tests:

// calculateRepetitions:
// s2_advance = 0
// total_count = 0
// for each letter_s1 in S1:
//  advance by one symbol
//  if s2_advance >= len_s2:
//    total_count +=1
//    s2_advance = 0
//    continue
//  if letter_s1 == s2[s2_advance]:
//    s2_advance+=1
//
// will calculate repetitions of s2 within s1

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if len(s1) == 0 || len(s2) == 0 {
		return -1
	}

	s2_advance := 0
	total_count := 0
	s2_runes := []rune(s2)
	s1_runes := []rune(s1)
	len_s2 := len(s2_runes)
	len_s1 := len(s1_runes)
	s1_size := len_s1 * n1
	s2_size := len_s2 * n2

	current_letter := s2_runes[0]

	// iterate over each letter in a first string
	for i := 0; i < s1_size; i++ {
		letter := s1_runes[i%len_s1]

		// skip cycle (letter does not match)
		if letter != current_letter {
			continue
		}

		// letters match, advance matcher to the next letter
		s2_advance += 1

		// current match is within bounds
		if s2_advance < s2_size {
			current_letter = s2_runes[s2_advance%len_s2]
			continue
		}
		// otherwise, reset counter to 0
		current_letter = s2_runes[0]
		s2_advance = 0
		total_count += 1
	}
	if s2_advance >= s2_size {
		total_count += 1
	}
	return total_count
}
