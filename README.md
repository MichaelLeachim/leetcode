# LeetLog 
TODO (SQL):

* https://leetcode.com/problems/human-traffic-of-stadium/
* https://leetcode.com/problems/trips-and-users/
* https://leetcode.com/problems/department-top-three-salaries/
* https://leetcode.com/problems/nth-highest-salary/

TODO: 
* [10 minute problem, for starting out] https://leetcode.com/problems/number-of-1-bits/
* https://leetcode.com/problems/recover-binary-search-tree/  
* https://leetcode.com/problems/sum-of-distances-in-tree/
* https://leetcode.com/problems/substring-with-concatenation-of-all-words/
* https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
* https://leetcode.com/problems/numbers-with-repeated-digits/


This is a Leet log. A log of Leetcode problems and their solutions done by me, </br>
in reverse chronological order. </br>

**Useful commands**

```shell
go test
go test -run TestMinCameraCover
go test -bench BenchmarkMinCameraCover
```

</br>

* **[2019-03-30]** [Talles billboard](https://leetcode.com/problems/tallest-billboard/)</br>
  ~10 hours and more, three days,  and five approaches later,  finally, solved it only to be in the bottom 5% of efficiency!</br>
  And this is for topic that I somewhat already familiar with. </br>
  One year won't be enough for me to prepare for an interview. </br>
  But, nonetheless, the solution was accepted, so I move on:</br>
  Submission details: https://leetcode.com/submissions/detail/218703638/</br>
  
* **[2019-03-27]** [Minimum window substring](https://leetcode.com/problems/minimum-window-substring/)</br>
  Well, this one was hard AF (well, as almost all problems submitted by leetcode).</br>
  The most difficult thing was to figure out how to check for pattern incrementally and fast</br>
  For the pattern with the unique letters the counting problem is trivial. But, for pattern with repeating letters it is hard</br>
  The most annoying and difficult thing is to figure out and debug this stuff. </br>
  On the plus side, it is the first my  solution where I am faster than 78% of submitted solutions. </br>
  And this freaking thing got 268 test cases! </br>
  In the end, I've given up debugging 60 lines of a total mess and wrote a nice object oriented, decomplected iterator and tested each part of </br>
  functionality seprately </br>
  So, it may be worth a look. </br>
  Anyway, here are submission details:</br>
  Submission details: https://leetcode.com/submissions/detail/218058420/</br>

* **[2019-03-25]** [Maximal rectangle](https://leetcode.com/problems/maximal-rectangle/)</br>
  It is the first problem, which I solved by iteration, not by recursion. </br>
  I am still in the bottom 25% by time efficiency. So, here is that</br>
  But, at least, I am learning now how to make it work. Overall, solving this problem </br>
  took me ~3 hours of deep thinking</br>
  I am not reading top </br>
  Submission details: https://leetcode.com/submissions/detail/217529915/

* **[2019-03-18]** [Binary Tree Cameras](https://leetcode.com/problems/binary-tree-cameras/)</br>
  Was sick this week, so, didn't add solutions. Now I am continuing working on these things.  </br>
  It took me ~3 hours to think up and code a solution for this problem. </br>
  Still, I am in the bottom of submissions (64 ms against 6 ms). But, at least, I have this thing solved. </br>
  Later, when I solve all of the hard ones, I will work on optimizations. 
  Submission details: https://leetcode.com/submissions/detail/215779712/
  
* **[2019-03-12]** [Odd even jump](https://leetcode.com/problems/odd-even-jump/)</br>
  I am able to implement naive DP solutions. (Meaning I get the intuition behind the solution and </br>
  I can break down problem into subproblems to create a recursive definition) But I still cannot implement performant solutions </br>
  In particular, this problem requires using the **monotonic stack** technique, that I haven't found</br>
  in CPHB</br>
  So, after I solve first 100 of hard problems, I should work on the optimization of my solutions. </br>
  Submission details: https://leetcode.com/submissions/detail/214320084/</br>

* **[2019-03-11]** [The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/) </br>
  Overall, the idea is to not to spend time on tangential problems. For example, I've spent a lot of </br>
  time implementing `iterateBeginEndOverSkyLine`. Which I shouldn't have done. </br>
  I should have assumed that the average hard problem requires **no more than an hour** for </br>
  thinking and implementing. If I am thinking about implementing something longer, </br>
  then this solution is most probably not the optimal one</br>
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
  

