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
    dayOrder = dict(zip(self.WEEKDAYS,range(0,7)))
    
    s = self.opening_hours
    if len(s) == 0:
      return ""
    
    first,last,hours = s[0][0],s[0][0],s[0][1]
    solution = {hours:[first,last,hours]}
    for day,hours in s[1:]:
      lastInSolution = len(solution)-1
      if hours in solution:
        (prevDayFirst,prevDayLast, prevHours) = solution[hours]
        solution[hours] = [prevDayFirst,day, hours]
        continue
      solution[hours] = [day,day,hours]
      
    preparedSolution = sorted(solution.values(),key=lambda x: dayOrder[x[0]])
    
    return ", ".join([firstDay + ": " + hours if firstDay == lastDay else firstDay + " - " + lastDay + ": " + hours for [firstDay,lastDay,hours] in preparedSolution])
  
class OpeningHour():
  def __init__(self, opening_hour, closing_hour):
    self.opening_hour = opening_hour
    self.closing_hour = closing_hour
  def to_string(self):
    return "{}-{}".format(self.opening_hour, self.closing_hour)
