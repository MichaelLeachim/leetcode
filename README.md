# LeetLog 

TODO: 
* [10 minute problem, for starting out] https://leetcode.com/problems/number-of-1-bits/
* https://leetcode.com/problems/recover-binary-search-tree/  
* https://leetcode.com/problems/sum-of-distances-in-tree/
* https://leetcode.com/problems/binary-tree-cameras/
* https://leetcode.com/problems/odd-even-jump/
* https://leetcode.com/problems/substring-with-concatenation-of-all-words/
* https://leetcode.com/problems/maximal-rectangle/
* https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

This is a Leet log. A log of Leetcode problems and their solutions done by me, </br>
in reverse chronological order. </br>

</br>
* **[2019-03-11]** [218. The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/)</br>
  One thing to consider:
    * Should not implement separate logic for ending calculations
  Overall, the idea is to not to spend time on tangential problems. For example, I've spent a lot of </br>
  time implementing `iterateBeginEndOverSkyLine`. Which I shouldn't have done. </br>
  I should assume that the average hard problem requires no more than an hour for </br>
  thinking and implementing. If I am thinking about implementing something longer, </br>
  then this solution is not the optimal one</br>
  Submission details: https://leetcode.com/submissions/detail/213984432/</br>

* **[2019-03-09]**  [Largest Component Size by Common Factor](https://leetcode.com/problems/largest-component-size-by-common-factor/)
  Solved this problem by using **union find** structure. Good thing about this problem is that it had me to study:
    * what is a **component** of a graph
    * what is a **factorization** of a number (for example, I didn't know that 1 wasn't prime, and didn't know any factorization algorithms)
    * what is a **union find** structure. I haven't heard about it before 
  Overall, although, I spent more time on it. (I should solve one/day, not one/2days), this problem was beneficial to me.  
  Submission details: https://leetcode.com/submissions/detail/213483591/
  
* **[2019-03-07]**  [Largest rectangle in histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)</br>
  That wasn't easy at all. First, I tried dynamic programming approach. It passed tests except tle on large input.</br>
  Then, I tried stack based solution. (That I yesterday decided to make). It didn't work either. Today, I tried another</br>
  way to handle this using dynamic programming reducing input through removing one of the params to the DP solve function.</br>
  Then, translated that algorithm into stack based. </br>
  And what is in the end?  My solution in the **bottom 10% of efficiency**. I guess, if I've just calculated all possible solutions</br>
  and choose the best it would work faster. </br>
  Anyways, at least, the solution was accepted, so I move on. </br>
  Later, when I will solve first 100 or so of these problems. I will return back and try to optimize it</br>
  [TODO]: (implement solution using segment trees, study Binary indexed tree)</br>
  Submission details: https://leetcode.com/submissions/detail/213024073/</br>
  
* **[2019-03-05]** [Basic calculator](https://leetcode.com/problems/basic-calculator/)</br>
  Implemented tokenizer and calculator for this problem </br>
  For possible optimization, should make streaming tokenizer, not consume entire input into memory</br>
  Submission details: https://leetcode.com/submissions/detail/212534416/

* **[2019-03-05]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Iterative solution is now working. I've made it by emulating `call stack` and using `goto` statements to </br>
  make jumps to recursive points. </br>
  Submission details: https://leetcode.com/submissions/detail/212493978/</br>
  
* **[2019-03-04]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Iterative solution: Ok, that wasn't simple. Tomorrow, implement conditional jumps in iterative algorithm.
  Emulate call stack. Figure out how to make return addresses/jumps within the subroutine. 
  (I am too tired today already)
  
* **[2019-03-04]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Recursive solution: https://leetcode.com/submissions/detail/212317430/
  
* **[2019-03-04]**   [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)</br>
  Took me 40-50 minutes to solve it. The trick to do it in one pass, is to iterate first to the `N`, and then walk forward </br>
  maintaining two pointers: one for current, one for `current+N`, when `N is nil`, `current.next = current.next.next` </br>
  Submission details: https://leetcode.com/submissions/detail/212314347/ </br>
  
* **[2019-03-04]**   [472. Concatenated Words](https://leetcode.com/problems/concatenated-words/) </br>
  Took me long time. Because was implementing trie from scratch </br>
  In the end, replaced trie with a simple hashmap. Ideas to use: </br>
  * first, maximally formalize the problem 
  * Second, write the code and tests
  * Third, **do not overoptimize**. Use whatever structures you already have in stdlib. Trie took me ~4 hours 
    to work and pass all tests and still I am not sure whether it is an optimal trie, or there are more 
    tricks to it.
  Submission details: https://leetcode.com/submissions/detail/212304244/
  
* **[2019-03-02]**   [Count the repetitions](https://leetcode.com/problems/count-the-repetitions/) </br>
  Done for ~3 hours to pass Leetcode test</br>
  Submission details: https://leetcode.com/submissions/detail/211732286/ </br>
* **[2019-03-01]**   [Cut off trees for golf](https://leetcode.com/problems/cut-off-trees-for-golf-event/). </br>
  Done for 2.5 hours to pass Leetcode test. </br>
  Submission details: https://leetcode.com/submissions/detail/211533228/ </br>
* **[2019-02-28]**  [Word break II](https://leetcode.com/problems/word-break-ii/). </br>
  Done for 4 hours to pass Leetcode test. </br>
  Submission details: https://leetcode.com/submissions/detail/211328790/ </br>
  

