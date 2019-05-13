# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-05-13 10:03 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

from symmetric_tree import Solution,TreeNode

def testIsPalindrome():
  assert Solution().isPalindrome([1,2,3,4,5,4,3,2,1]) == True
  assert Solution().isPalindrome([1,2,3,4,5,5,3,2,1]) == False
  assert Solution().isPalindrome([]) == True
  assert Solution().isPalindrome([1]) == True
  assert Solution().isPalindrome([1,2]) == False
  assert Solution().isPalindrome([1,2,1]) == True
  assert Solution().isPalindrome([1,1]) == True
  assert Solution().isPalindrome([1,2]) == False
  assert Solution().isPalindrome([1,2,2,1]) == True
  assert Solution().isPalindrome([1,2,2,2]) == False
  
  
def testSymmetricTree():
  root = TreeNode(1)
  root.left = TreeNode(2)
  root.left.left = TreeNode(3)
  root.left.right = TreeNode(4)
  
  root.right = TreeNode(2)
  root.right.left = TreeNode(4)
  root.right.right = TreeNode(3)
  
  result = []
  def resfn(x,y):
    result.append(x)
    return True

  Solution().bfs(root, resfn)
  assert result == [1, 2, 2, 3, 4, 4, 3]
  
  
  assert 1 == 1
  pass 
