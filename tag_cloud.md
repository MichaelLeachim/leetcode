# The idea is to use segment tree for the tag cloud

How exactly can we use that? 

Example, nodes list:

1 hard_problems, work, home
2 easy_problems, work, office
3 medium_problems, hobby, home
4 thinking, programming
5 thinking, philosophy

hard_problems --> 1
easy_problems --> 2
work --> hard_problems
work --> easy_problems
home --> work
office --> work


