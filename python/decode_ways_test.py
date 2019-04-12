#!/usr/bin/python
# -*- coding: utf-8 -*-

# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-11 08:16 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

from decode_ways import Solution

def test_num_decodings():
  assert len(Solution().learningStuff()) == 26
  assert Solution().numDecodings("226") == 3
  
  assert Solution().numDecodings("12") == 2
  assert Solution().numDecodings("0") == 0
  assert Solution().numDecodings("0123") == 0
  
  assert Solution().numDecodings("10") == 1
  
  assert Solution().numDecodings("") == 0
  
  assert Solution().numDecodings("1111111") == 21
  assert Solution().numDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253") == 3981312
  
# def test_num_decodings_bench(benchmark):
  
#   benchmark(lambda: Solution().numDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253"))
  
