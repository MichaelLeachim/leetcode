# LeetLog 

TODO: 
* [10 minute problem, for starting out] https://leetcode.com/problems/number-of-1-bits/
* https://leetcode.com/problems/recover-binary-search-tree/  
* https://leetcode.com/problems/sum-of-distances-in-tree/
* https://leetcode.com/problems/binary-tree-cameras/

This is a Leet log. A log of Leetcode problems and their solutions done by me, </br>
in reverse chronological order. </br>
</br>
* **[2019-03-05]** [Basic calculator](https://leetcode.com/problems/basic-calculator/)</br>
  Implemented tokenizer and calculator for this problem (Don't know why it is in hard section, but solved it nonetheless)</br>
  Submission details: https://leetcode.com/submissions/detail/212534416/

* **[2019-03-05]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Iterative solution is now working. I've made it by emulating `call stack` and using `goto` statements to </br>
  make jumps to recursive points. </br>
  Submission details: https://leetcode.com/submissions/detail/212493978/</br>
  
* **[2019-03-04]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Iterative solution: Ok, that wasn't simple. Tomorrow, implement conditional jumps in iterative algorithm.
  Emulate call stack. Figure out how to make return addresses/jumps within the subroutine. 
  (Nothing extraordinary, but I am too tired today already)
  
* **[2019-03-04]** [PostOrderTraversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)</br>
  Recursive solution: https://leetcode.com/submissions/detail/212317430/
* **[2019-03-04]**   [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)</br>
  Trivial problem and of medium difficulty, but I wanted it to solve because I've never really worked with linked lists. </br>
  Took me 40-50 minutes to solve it. The trick to do it in one pass, is to iterate first to the `N`, and then walk forward </br>
  maintaining two pointers: one for current, one for `current+N`, when `N is nil`, `current.next = current.next.next` </br>
  Submission details: https://leetcode.com/submissions/detail/212314347/ </br>
  
  
* **[2019-03-04]**   [472. Concatenated Words](https://leetcode.com/problems/concatenated-words/) </br>
  Took me long time. Because was implementing trie from scratch, which is a lot of boring work :( 
  In the end, replaced trie with a simple hashmap. Tricks to use: 
  * first, maximally formalize the problem 
  * Second, write the code and tests
  * Third, **do not overoptimize**. Use whatever structures you already have in stdlib. Trie took me ~4 hours 
    to work and pass all tests and still I am not sure whether it is an optimal trie, or there are more 
    tricks to it.
  * And the last. Field related topics on Leetcode lies. 
  
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
  

