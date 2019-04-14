# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
# @ Copyright (c) Michael Leahcim                                                      @
# @ You can find additional information regarding licensing of this work in LICENSE.md @
# @ You must not remove this notice, or any other, from this software.                 @
# @ All rights reserved.                                                               @
# @@@@@@ At 2019-04-14 08:19 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
from collections import OrderedDict
import ipdb

# Okay, how to solve it?

# For every item, if previous hours match current hours,
# Then, merge previous item with current item
# Otherwise add them to array.

class Restaurant():
  WEEKDAYS = [
   'Sun',
   'Mon',
   'Tue',
   'Wed',
   'Thu',
   'Fri',
   'Sat',
  ]
  def __init__(self, *opening_hours):
    # group day with the opening hours
    # [('Mon', '9-17'), ('Tue', '10-15'), ...]
    self.opening_hours = zip(
        self.WEEKDAYS,
        [opening_hour.to_string() for opening_hour in opening_hours])
  def get_opening_hours(self):
    # ipdb.set_trace()
    s = self.opening_hours
    solution = s[0]
    
    for index, [day,hours] in enumerate(s[1:]):
      index = index+1
      [prevday, prevhours] = solution[index-1]
      if prevhours == hours:
        prevday = prevday + " - " + day
        solution[index-1] = [prevday,hours]
        continue
      solution.append(day,hours)
          
    
    

    
    
    return ""
  
class OpeningHour():
  def __init__(self, opening_hour, closing_hour):
    self.opening_hour = opening_hour
    self.closing_hour = closing_hour
  def to_string(self):
    return "{}-{}".format(self.opening_hour, self.closing_hour)
