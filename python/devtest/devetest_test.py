# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-14 08:17 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
from devtest import Restaurant, OpeningHour
import unittest

class RestaurantTestCase(unittest.TestCase):
  def test_empty(self):
    restaurant = Restaurant()
    self.assertTrue(len(restaurant.get_opening_hours()) == 0)
  def test_simple(self):
    restaurant = Restaurant(
      OpeningHour(8, 16),  # Sunday
      OpeningHour(8, 17),  # Monday
      OpeningHour(8, 18),  # Tuesday
      OpeningHour(8, 19),  # Wednesday
      OpeningHour(8, 20),  # Thursday
      OpeningHour(8, 21),  # Friday
      OpeningHour(8, 22),  # Saturday
    )
    self.assertEqual(
      "Sun: 8-16, Mon: 8-17, Tue: 8-18, Wed: 8-19, Thu: 8-20, Fri: 8-21, Sat: 8-22",
      restaurant.get_opening_hours())
    
  def test_single_group(self):
    restaurant = Restaurant(
      OpeningHour(8, 16),  # Sunday
      OpeningHour(8, 16),  # Monday
      OpeningHour(8, 16),  # Tuesday
      OpeningHour(8, 16),  # Wednesday
      OpeningHour(8, 20),  # Thursday
      OpeningHour(8, 21),  # Friday
      OpeningHour(8, 22),  # Saturday
    )
    self.assertEqual(
      "Sun - Wed: 8-16, Thu: 8-20, Fri: 8-21, Sat: 8-22",
      restaurant.get_opening_hours())
    
  def test_merge_test(self):
    restaurant = Restaurant(
      OpeningHour(8, 15),  # Sunday
      OpeningHour(8, 16),  # Monday
      OpeningHour(8, 16),  # Tuesday
      OpeningHour(8, 16),  # Wednesday
      OpeningHour(8, 16),  # Thursday
      OpeningHour(8, 16),  # Friday
      OpeningHour(8, 15),  # Saturday
    )
    mci = restaurant.mergeConsequentItems(restaurant.opening_hours)
    self.assertListEqual(mci,[['Sun', 'Sun', '8-15'], ['Mon', 'Fri', '8-16'], ['Sat', 'Sat', '8-15']])
    mnc = restaurant.mergeNonConsequentItems(mci)
    self.assertListEqual(mnc,[[['Sun', 'Sun', '8-15'], ['Sat', 'Sat', '8-15']], [['Mon', 'Fri', '8-16']]])
    
  def test_multiple_groups(self):
      restaurant = Restaurant(
          OpeningHour(8, 16),  # Sunday
          OpeningHour(8, 16),  # Monday
          OpeningHour(8, 16),  # Tuesday
          OpeningHour(8, 17),  # Wednesday
          OpeningHour(8, 18),  # Thursday
          OpeningHour(8, 20),  # Friday
          OpeningHour(8, 20),  # Saturday
      )
      self.assertEqual(
          "Sun - Tue: 8-16, Wed: 8-17, Thu: 8-18, Fri - Sat: 8-20",
          restaurant.get_opening_hours())
      
  def test_edge_case(self):
    restaurant = Restaurant(
      OpeningHour(8, 16),  # Sunday
      OpeningHour(8, 17),  # Monday
      OpeningHour(8, 17),  # Tuesday
      OpeningHour(8, 17),  # Wednesday
      OpeningHour(8, 16),  # Thursday
      OpeningHour(8, 16),  # Friday
      OpeningHour(8, 16),  # Saturday
    )
    self.assertEqual(
      "Sun, Thu - Sat: 8-16, Mon - Wed: 8-17",
      restaurant.get_opening_hours())
    
  def test_edge_case2(self):
    restaurant = Restaurant(
      OpeningHour(8, 15),  # Sunday
      OpeningHour(8, 16),  # Monday
      OpeningHour(8, 16),  # Tuesday
      OpeningHour(8, 16),  # Wednesday
      OpeningHour(8, 16),  # Thursday
      OpeningHour(8, 16),  # Friday
      OpeningHour(8, 15),  # Saturday
    )
    self.assertEqual(
      "Sun, Sat: 8-15, Mon - Fri: 8-16",
      restaurant.get_opening_hours())
