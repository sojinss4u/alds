import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
Q: Check if a given string is paliandrome
A string is paliandrome if it reads the same backward & forward
Time Complexity: O(n) [Because of For Loop]
Space Complexity: O(1) [No increase in memory when string size increase]
"""

def is_paliandrome_without_using_inbuilt_function(s):
    reversed_string = ""
    for index in reversed(range(len(str(s)))):
        reversed_string = reversed_string + str(s)[index]
    if str(s) == reversed_string:
        return True
    else:
        return False


def is_paliandrome_using_inbuilt_string_reverse(s):
   reverse_string = str(s)[::-1]
   if str(s) == reverse_string:
     return True
   else:
     return False

if __name__ == "__main__":
  s = 1211
  result = is_paliandrome_without_using_inbuilt_function(s)
  logger.info(result)
