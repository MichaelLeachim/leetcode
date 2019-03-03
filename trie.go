// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-03 21:44 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

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
				trie[cur_pos][letter] = len(trie)
				trie = append(trie, [26]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1})
			}
			cur_pos = trie[cur_pos][letter]
		}
		cur_pos = 0
	}
	return trie
}

func traverseAndQueryTrie(query string, trie [][26]int) []string {
	curNode := 0
	for _, letter := range query {
		curNode = trie[curNode][letter%26]
		if curNode == -1 {
			return []string{}
		}
	}
	result := []string{}
	for _, word := range queryTrie(curNode, trie) {
		result = append(result, query+word)
	}
	if len(result) == 0 {
		result = append(result, query)
	}
	return result
}

func queryTrie(startingNode int, trie [][26]int) []string {
	// queryTree(0) = letter+queryTrie(loc) for letter,loc in trie[0], if loc exists.
	//                if len(queryTrie(loc)) == 0, letter

	var dfs func(s int) [][]int
	// perform dfs search to find all suffixes
	dfs = func(curNode int) [][]int {
		wordlist := [][]int{}
		if curNode >= len(trie) || curNode < 0 {
			return wordlist
		}

		// for letter,loc in trie[N]
		for letter, loc := range trie[curNode] {
			// if loc exists
			if loc != -1 {
				dfs_realized := dfs(loc)
				if len(dfs_realized) == 0 {
					wordlist = append(wordlist, []int{letter})
					continue
				}
				for _, suffix := range dfs_realized {
					wordlist = append(wordlist, append([]int{letter}, suffix...))
				}
			}
		}
		return wordlist
	}

	// return human readable strings
	result := []string{}
	for _, word := range dfs(startingNode) {
		row := []rune{}
		for _, letter := range word {
			row = append(row, ascii_table[letter])
		}
		result = append(result, string(row))
	}
	return result
}
