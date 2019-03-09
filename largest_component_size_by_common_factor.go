// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-08 10:17 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

// Solving https://leetcode.com/problems/largest-component-size-by-common-factor/ problem
// [4,6,15,35]
// {4,6,15,35}
// [20,50,9,63]
// {20,50} {9,63}
// [2,3,6,7,4,12,21,39]
// {2,3,6} {7}, {4,12,21,39}
// {2,3,6} {4,7,12,21,39}
// {2,3,4,6,7,12,21,39}
// {2,3,4,6,7,21,12,39}
// {2,6,4,12} {3,21,39} {7} {13}

// {2,3,4,6,7,12,21,39}
// {2,3,4,6,7,12,21,39}

// {2}
// {3}
// {2}

// {2,3,7,13},

//  [2,3,6,7,4,12,21,39]
//  {2,3} {7}

// The algorithm:
//  For each element in N:
//    representatives := {}
//    for each factor in factorize(element):
//       add its representative into representatives
//       if factor is within set, add  its representative to represenatives set
//    join representatives and increment best counter

func largestComponentSize(A []int) int {
	link, size, numcount := map[int]int{}, map[int]int{}, map[int]int{}
	best := 0

	// gets the representative of the element
	find := func(x int) int {
		for x != link[x] {
			x = link[x]
		}
		return x
	}

	// unites two sparse sets
	unite := func(a, b int) {
		a, b = find(a), find(b)
		if size[a] < size[b] {
			b, a = a, b
		}
		size[a] += size[b]
		numcount[a] += numcount[b]
		if numcount[a] > best {
			best = numcount[a]
		}

		link[b] = a
	}

	// factorizes integer. Returns all prime divisors of a given number
	factorize := func(a int) []int {
		result := []int{}
		for i := 2; i*i <= a; i++ {
			for a%i == 0 {
				result = append(result, i)
				// TODO, why is that?
				a /= i
			}
		}
		if a > 1 {
			result = append(result, a)
		}
		return result
	}

	for _, item := range A {
		undub := map[int]bool{}
		// get representatives, remove duplicates
		representatives := []int{}

		for _, factor := range factorize(item) {
			// does this factor exist?
			_, ok := link[factor]
			if !ok {
				link[factor] = factor
				size[factor] = 1
				representatives = append(representatives, factor)
				undub[factor] = true
				continue
			}

			repr := find(factor)
			// if duplicate, continue
			if undub[repr] {
				continue
			}
			// track this number in numcount
			undub[repr] = true
			representatives = append(representatives, repr)
		}
		if len(representatives) == 0 {
			continue
		}
		first, rest := representatives[0], representatives[1:]
		// join up representatives together
		for _, repr := range rest {
			unite(first, repr)
		}

		nf := find(first)
		ncf := numcount[nf] + 1
		if ncf > best {
			best = ncf
		}
		numcount[nf] = ncf

	}
	return best
}
