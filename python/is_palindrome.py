# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-23 20:44 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
class Palindrome:
  @staticmethod
  def is_palindrome(word):
    if not isinstance(word,str):
      return False
    #Please write your code here.
    is_palindrome = True
    word = word.lower()
    for right,left in zip(xrange(0,len(word)),xrange(len(word)-1,-1,-1)):
      if right == left:
        return is_palindrome
      if word[right] != word[left]:
        return False
    return is_palindrome
