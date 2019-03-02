// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Problem: https://leetcode.com/problems/concatenated-words/

// Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

// Compound words are always longer than simple words that they are part of
// sorted(x)
// Solution(x) returns ([simple words of x], [compound words of x])
// Solution(x) = Solution(x-1) + if(x in simple words) => do nothing
//                                                     => check against simple words,
//                                                          if match, then to [concatenaded words]
//                                                          otherwise      to [simple words]

// check(A,x) =
//  if A or x is empty: false
//  if x[len(x)-1] is a substring of A, then check(A without x[len(x)-1],x-1) if (x-1 and (A without  x[len(x)-1])) are not empty

//  Tricky usecase for greedy check: ["catsdog","cat","sdog","dog"], if match "dog", then the "catsdog" is not going to be matched
// check("catdog") => 1
// check("catdo") =>  0
// check("catd")  =>  0
// check("cat")   =>  1
// check(X) =>
// for suffix in matchAllSuffixes(X):
//   if(suffix == X): return true (found match)
//   if check(X without suffix):
//     return true
// return false

// constructTrie
// according to the problem definition: the maximum N is 600000, 26 is the number
// of letters in English

var ascii_table = (func() [26]rune {
	a := [26]rune{}
	a[int('a')%26] = 'a'
	a[int('b')%26] = 'b'
	a[int('c')%26] = 'c'
	a[int('d')%26] = 'd'
	a[int('e')%26] = 'e'
	a[int('f')%26] = 'f'
	a[int('g')%26] = 'g'
	a[int('h')%26] = 'h'
	a[int('i')%26] = 'i'
	a[int('j')%26] = 'j'
	a[int('k')%26] = 'k'
	a[int('l')%26] = 'l'
	a[int('m')%26] = 'm'
	a[int('n')%26] = 'n'
	a[int('o')%26] = 'o'
	a[int('p')%26] = 'p'
	a[int('q')%26] = 'q'
	a[int('r')%26] = 'r'
	a[int('s')%26] = 's'
	a[int('t')%26] = 't'
	a[int('u')%26] = 'u'
	a[int('v')%26] = 'v'
	a[int('w')%26] = 'w'
	a[int('x')%26] = 'x'
	a[int('y')%26] = 'y'
	a[int('z')%26] = 'z'
	return a
})()

func constructTrie(words []string) [][26]int {
	trie := [][26]int{[26]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}
	cur_pos := 0
	for _, word := range words {
		for _, letter := range word {
			letter = letter % 26
			// in case, no next node
			if trie[cur_pos][letter] == -1 {
				trie[cur_pos][letter] = cur_pos + 1
				trie = append(trie, [26]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1})
			}
			cur_pos = trie[cur_pos][letter]
		}
		cur_pos = 0
	}
	return trie
}

func queryTrie(startingNode int, trie [][26]int) []string {
	var dfs func(s, e int) [][]int
	// // drill down for the given query string
	// curNode := trie[0][query[0]%26+1]
	// for _, letter := range query[1:] {
	// 	curNode = trie[curNode][letter%26+1]
	// }

	// perform dfs search to find all suffixes
	dfs = func(curNode, oldNode int) [][]int {
		wordlist := [][]int{}
		// for each letter
		for letter, child := range trie[curNode] {
			if child != -1 && child != oldNode {

				dfs_realized := dfs(child, curNode)
				for _, suffix := range dfs_realized {
					wordlist = append(wordlist, append([]int{letter}, suffix...))
				}
				if len(dfs_realized) == 0 {
					wordlist = append(wordlist, []int{letter})
				}
			}

		}

		return wordlist
	}
	// return human readable strings
	result := []string{}
	for _, word := range dfs(startingNode, -1) {
		row := []rune{}
		for _, letter := range word {
			row = append(row, ascii_table[letter])
		}
		result = append(result, string(row))
	}
	return result
}

func findAllConcatenatedWordsInADict(words []string) []string {
	return []string{}
}
