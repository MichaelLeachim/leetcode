#!/usr/bin/python
# -*- coding: utf-8 -*-

# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-11 08:16 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

from decode_ways import Solution

def test_dict_list_invariant_of_solution():
  sols = ['1', '11','111','111','111','111','11','131313', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1']
  newsols = []
  for item in sols:
    if item != '1' and item != '11':
      newsols.append(item)
  assert newsols == ['111', '111', '111', '111', '131313']
  ## and its invariant in compressed form
  ### first, compress the input
  sols2 = {}
  for item in sols:
    if item in sols2:
      sols2[item]+=1
      continue
    sols2[item] = 1
    
  newsols2 = {}
  for item,count in sols2.iteritems():
    if item != '1' and item != '11':
      newsols2[item] = count
  assert  {'111':4,'131313':1}   == newsols2
  assert len(sols) == sum(sols2.values())

  something = ['1', '11', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1']
  somethingDict = {}
  for item in something:
    if item in somethingDict:
      somethingDict[item]+=1
      continue
    somethingDict[item] = 1
    
  assert len(something) == sum(somethingDict.values())
  i = {"a":3}
  Solution().sumOrSet(i,"a",5)
  assert i == {"a":8}

def test_dictionary_workage():
  printsol = {}
  previousSolutions = ['1', '11', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1', '1', '11', '1', '11', '1', '1', '11', '1']
  for i in previousSolutions:
    if i in printsol:
      printsol[i]+=1
    else:  
      printsol[i] = 1
  assert printsol == {'1':13,'11':8}  

def test_num_decodings():
  
  assert len(Solution().learningStuff()) == 26
  assert Solution().numDecodings("226") == 3
  
  assert Solution().numDecodings("12") == 2
  assert Solution().numDecodings("0") == 0
  assert Solution().numDecodings("0123") == 0
  
  assert Solution().numDecodings("10") == 1
  
  assert Solution().numDecodings("") == 0
  
  assert Solution().numDecodings("1111111") == 21
  
  # assert Solution().numDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253") == 3981312
  
# def test_num_decodings_bench(benchmark):
  
#   benchmark(lambda: Solution().numDecodings("9371597631128776948387197132267188677349946742344217846154932859125134924241649584251978418763151253"))
  
