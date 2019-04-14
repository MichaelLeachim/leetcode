# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-13 19:25 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

# Solving: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
# [3,3,5,0,0,3,1,4]

## Calculate difference matrix

## Difference matrix

## matrix[buyDay][sellDay] = profit

## What is the new problem definition?
### Find such buyDay and sellDay and  buyDay2, sellDay2
###   where their sum is maximum
###   and, buyDay2 >= sellDay

### Approaches?
#### We can find the best solution on the left and on the right for every X_n in input
####   Just iterate over it, and calculate the sum of the given input, choose the best

#### Now, how we can find the best solution for up to N?
###### Can we take the best solution for N-1?
###### Input is: [1,3,2,4]
######   Solution(1)   =   = max(solution(),(1-1))
######   Solution(1,3)     = max(solution(1),3-1,3-3)
######   Solution(1,3,2)   = max(solution(1,3),2-1,2-3,2-2)
######   Solution(1,3,2,4) = max(solution(1,3,2),4-1,4-3,4-2,4-4)

#### Now, we should repeat this process backwards, to find the best solution on the right
#    side of the input
#    Solution(4)   = max(solution(),(4-4))
#    Solution(2,4) = max(solution(4),2-2,4-2)
#    Solution(3,2,4) = max(solution(2,4),3-3,2-3,4-3)
#    Solution(1,3,2,4) = max(solution(3,2,4),1-1,3-1,2-1,4-1)

#### Now, how to optimize?
# I don't like repeated subtractions. On large data size it will take a lot of time
# We can infer subtractions by previous, like this:
##   Solution(1)   =   = max(solution(),(1-1))
##   Solution(1,3)     = max(solution(1),(1-1)+(3-1),3-3)
##   Solution(1,3,2)   = max(solution(1,3),(1-1)+(3-1)+(2-3),(3-3)+(2-3),2-2)
##   Solution(1,3,2,4) = e.t.c.

### No no no, there is an easier way
##  If we find min from input for every step from left and right?
##    [1,3,2,4] => [1,1,1,1]
##  And then, we just subtract that minimum from every point

class Solution(object):
  def calculateMatrix(self,input):
    return [[item2-item if index2>=index else None for index2,item2 in enumerate(input)]
            for index,item in enumerate(input)]
  
  def calculateBestForward(self,input):
    prev     = 0
    yield prev
    
    if len(input) == 0:
      return
    minimum  = input[0]
    
    for index,item in enumerate(input[1:]):
      index = index+1
      prev = max(item-minimum,prev)
      if minimum > item:
        minimum = item
      yield prev
      
  def calculateBestBackward(self,input):
    prev = 0
    yield prev

    if len(input) == 0:
      return
    maximum  = input[len(input)-1]
    
    for index in xrange(len(input)-2,-1,-1):
      item = input[index]
      prev = max(maximum-item,prev)
      if item > maximum:
        maximum = item
      yield prev
  
  def maxProfit(self, prices):
    """
    :type prices: List[int]
    :rtype: int
    """
    best = 0
    for forward,backward in zip(self.calculateBestForward(prices),reversed(list(self.calculateBestBackward(prices)))):
      if forward+backward > best:
        best = forward+backward
    return best      

