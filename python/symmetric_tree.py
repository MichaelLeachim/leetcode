# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-05-13 09:59 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

## Solving: https://leetcode.com/problems/symmetric-tree/

from collections import deque


# BFS on the tree:
#  For every tree level, check if palindrome, if not, return false

# Definition for a binary tree node.

class TreeNode(object):
  def __init__(self, x):
    self.val = x
    self.left = None
    self.right = None
    
class Solution(object):
  
  def isPalindrome(self,input):
    inlen = len(input)
    
    for index,val in enumerate(input[:inlen/2]):
      if input[(inlen-index)-1] != val:
        return False
    return True
  
  def bfs(self,root,processfn):
    q = deque([root])
    visited = {root:True}
    distance = {root:0}
    # import ipdb; ipdb.set_trace()
    while q:
      s = q.pop()
      res = processfn(s.val,distance[s])
      if not res:
        return
      
      for node in [s.left,s.right]:
        if (not node) or (node in visited):
          continue
        visited[node] = True
        distance[node] = distance[s]+1
        q.appendleft(node)
        
  def isSymmetric(self, root):
    """
    :type root: TreeNode
    :rtype: bool
    """
    
    

    

        
      

      

    
    
    
    
    
    
    
    
    
        
