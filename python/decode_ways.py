# -*- coding: utf-8 -*-
#!/usr/bin/python

# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-11 08:12 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

# Solving Decode ways https://leetcode.com/problems/decode-ways/

# solution("2") => ["2"]
# solution("22") => ["22"], ["2","2"]
# solution("226") =>  ["226"], ["22", "6"], ["2","2","6"], ["2","26"]

# Or, if we simplify (take only the last number)
# solution("2") => "2"
# solution("22") => "22", "2"
# solution("226") => "226", "26", "6", "6"

# And, the same approach 
# solution("1") => ["1"]
# solution("12") => ["12"], ["1","2"]
# solution("123") => ["123"], ["1","23"], ["12","3"], ["1","2","3"]

# Or, simplified:
# solution("1") => "1"
# solution("12") => "12","2"
# solution("123") => "123","23", "3","3"

# The simplified formula is:
# solution(X) = for every solution(X-1), either concatenate, if possible, or append

## Possible optimization. Maintain dictionary with amounts of each entry instead of a flat list

class Solution(object):
  inputSet = {"1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26"}
  def learningStuff(self):
    return self.inputSet
  
  def numDecodings(self, s):
    """
    :type s: str
    :rtype: int
    """
    if len(s) == 0:
      return 0
    
    previousSolutions = {s[0]:1}
    if not previousSolutions.keys()[0] in self.inputSet:
      return 0
    
    for item in s[1:]:
      newPreviousSolutions = {}
      for solution,count in previousSolutions.iteritems():
        sol = "".join([solution,item])
        if sol in self.inputSet:
          newPreviousSolutions[sol] = count
          
      if item in self.inputSet:  
        newPreviousSolutions[item] = len(previousSolutions)
        
      previousSolutions = newPreviousSolutions
      print previousSolutions  
    
    return sum([item for item in previousSolutions.values()])
