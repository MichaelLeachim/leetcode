# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-13 19:28 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

from best_time_to_buy_and_sell_stock import Solution

def testMakeNiceDifferenceMatrix():
  N = None
  assert Solution().calculateMatrix([3,3,5,0,0,3,1,4]) == [
    [0, 0, 2, -3, -3, 0, -2, 1],
    [N, 0, 2, -3, -3, 0, -2, 1],
    [N, N, 0, -5, -5, -2, -4, -1],
    [N, N, N, 0, 0, 3, 1, 4],
    [N, N, N, N, 0, 3, 1, 4],
    [N, N, N, N, N, 0, -2, 1],
    [N, N, N, N, N, N, 0, 3],
    [N, N, N, N, N, N, N, 0]]
  
  assert Solution().calculateMatrix([1,3,2,4]) == [
    [0, 2,  1, 3],
    [N, 0, -1, 1],
    [N, N,  0, 2],
    [N, N,  N, 0]]
  
  mt = Solution().calculateMatrix([3,3,5,0,0,3,1,4])
  ## first is buyDay, second is sellDay. if sellDay is before buyDay, then return None
  assert mt[0][2] == 2
  assert mt[0][3] == -3
  assert mt[6][7] == 3
  assert mt[2][1] == None
  
def testCalculateBestForwardAndBackward():
  assert list(Solution().calculateBestForward([1, 3, 2, 4])) == [0,2,2,3]
  assert list(Solution().calculateBestForward([3,3,5,0,0,3,1,4])) == [0, 0, 2, 2, 2, 3, 3, 4]

  assert list(reversed(list(Solution().calculateBestBackward([1, 3, 2, 4])))) == [3,2,2,0]
  assert list(reversed(list(Solution().calculateBestBackward([3,3,5,0,0,3,1,4])))) == [4,4,4,4,4,3,3,0]

  
  
def testMaxProfit():
  assert Solution().maxProfit([3, 3, 5, 0, 0, 3, 1, 4]) == 6
  assert Solution().maxProfit([1, 3, 2, 4]) == 4
  assert Solution().maxProfit([1, 2, 3, 4, 5]) == 4
  assert Solution().maxProfit([7, 6, 4, 3, 1]) == 0
  assert Solution().maxProfit([]) == 0
  assert Solution().maxProfit([1]) == 0
  assert Solution().maxProfit([1,2]) == 1
  largeInput = range(10000,0)
  largeInput.extend([0]*100)
  assert Solution().maxProfit(largeInput) == 0

def testCalculateBestForwardBench(benchmark):
  largeInput = range(1000,0)
  largeInput.extend(range(0,1000))
  benchmark(lambda: Solution().calculateBestForward(largeInput))
  
def testCalculateBestBackwardBench(benchmark):
  largeInput = range(1000,0)
  largeInput.extend(range(0,1000))
  benchmark(lambda: Solution().calculateBestBackward(largeInput))
