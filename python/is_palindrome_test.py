# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-23 20:44 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
from is_palindrome import Palindrome

def testIsPalindrome():
  assert Palindrome.is_palindrome("Hell world") == False
  assert Palindrome.is_palindrome("") == True
  assert Palindrome.is_palindrome("aa") == True
  assert Palindrome.is_palindrome("aba") == True
  assert Palindrome.is_palindrome("abb") == False
  assert Palindrome.is_palindrome(None) == False
  assert Palindrome.is_palindrome("AaaAaa") == True

  


