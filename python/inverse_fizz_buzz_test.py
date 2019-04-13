# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-12 20:40 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

from inverse_fizz_buzz import InverseFizzBuzz
import unittest

class InverseFizzBuzzTestCase(unittest.TestCase):

  def testForward(self):
    inverse = InverseFizzBuzz(["fizz"])
    # self.assertDictEqual(inverse.forward,
    #                      {'fizzbuzz': [10, 20, 30, 40, 50, 60, 70, 80, 90, 100], 'fizz': [3, 6, 9, 12, 18, 21, 24, 27, 33, 36, 39, 42, 48, 51, 54, 57, 63, 66, 69, 72, 78, 81, 84, 87, 93, 96, 99], 'buzz': [5, 15, 25, 35, 45, 55, 65, 75, 85, 95]})


    
  def testFizz(self):
    inverse = InverseFizzBuzz(["fizz"])
    self.assertEqual([3], inverse.sequence())

  def testNone(self):
    inverse = InverseFizzBuzz(["buzz fizz buzz"])
    self.assertEqual(None, inverse.sequence())
    
  def testFizzBuzz(self):
    inverse = InverseFizzBuzz(["fizzbuzz"])
    self.assertEqual([10], inverse.sequence())

  def testFizzBuzzSimple(self):
    inverse = InverseFizzBuzz(["fizz","fizz"])
    self.assertEqual([6,7,8,9], inverse.sequence())
    
    
  def testFizzBuzzMore(self):
    inverse = InverseFizzBuzz(["fizzbuzz","fizzbuzz","fizzbuzz"])
    self.assertEqual(range(10,31), inverse.sequence())
    
  def testBuzz(self):
    inverse = InverseFizzBuzz(["buzz"])
    self.assertEqual([5], inverse.sequence())
    
  def testFizzBuzz(self):
    inverse = InverseFizzBuzz(["fizz", "buzz"])
    self.assertEqual([9, 10], inverse.sequence())
    
  def testBuzzFizz(self):
    inverse = InverseFizzBuzz(["buzz", "fizz"])
    self.assertEqual([5, 6], inverse.sequence())
    
  def testFizzBuzzFizz(self):
    inverse = InverseFizzBuzz(["fizz", "buzz", "fizz"])
    self.assertEqual([3, 4, 5, 6], inverse.sequence())
    
  def testFizzFizz(self):
    inverse = InverseFizzBuzz(["fizz", "fizz"])
    self.assertEqual([6, 7, 8, 9], inverse.sequence())
    
  def testFizzFizzBuzz(self):
    inverse = InverseFizzBuzz(["fizz", "fizz", "buzz"])
    self.assertEqual([6, 7, 8, 9, 10], inverse.sequence())
