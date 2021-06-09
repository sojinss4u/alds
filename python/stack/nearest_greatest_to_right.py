import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
Q: https://practice.geeksforgeeks.org/problems/next-larger-element-1587115620/1
   https://www.youtube.com/watch?v=NXOOYYwpbg4
Given an array ar[] of size 'n' having distinct elements, the task is to find the next greater element, for each element of the array in order of their apperance in the array.
Next greater element of an element, in the array is the nearest element on the right which is greater than the current element. If there does not exist next greater of the current element,
then next greater element for current element is -1. For example next greater for last element is always -1. 
"""

"""
Explanation:
  1. BruteForceMethod:
     In this method, we will take two nested for loops where the first loop variable will be 'i' & second loop variable will be 'j'
     result = []
     for i in range(len(ar)):
       for j in range(i+1,len(ar)):
         if ar[j] > ar[i]:
           result.append(ar[j])
           break
       if j == len(ar):
         result.append(-1)
    Time Complexity : O(n^2), due to nested for loops
    Space Complexity: O(n), due to storing result in result array

  2.  Optimized solution using stack
      As you can in the above brute force approach, 'j' is dependent on 'i'. Whenever 'j' is dependent on 'i', we can reduce the time complexity of the problem from O(n^2) to O(n) using a stack.
      In order to implement this solution using stack, we will have to start the first loop traversal from right to left. We will take each element in the array from right to left & then check 
      1. If s.top() > ar[i]
         If Yes, then the greatest element to right for the current element is stack top.
      2. If s.top() <= ar[i]
         Now we will pop() the top of the stack & check the next element. We will continue this untill we find a greater element in the stack
      3. If len(stack) == 0, then it means either we are at the last element or there are no greater element to the right of the current element
         Now we will return -1          
"""

class Stack:

    def __init__(self):
        self.list = list()

    def push(self,data):
        self.list.append(data)

    def pop(self):
        self.list.pop()

    def top(self):
        if self.length() > 0:
            return self.list[-1]
        else:
            return None

    def length(self):
        return len(self.list)

def find_next_greatest(ar):
    result = list()
    s = Stack()
    for item in reversed(ar):
        logger.debug(item)
        s_length = s.length()
        s_top = s.top()
        if s_length == 0:
            result.append(-1)
            s.push(item)
        elif s_top > item:
            result.append(s_top)
            s.push(item)
        else:
            while s_top < item and s_top != None:
                s.pop()
                s_top = s.top()
            if s_length == 0:
                result.append(-1)
            else:
                result.append(s_top)
            s.push(item)
    result.reverse()
    return result



if __name__ == "__main__":
    ar = [2,3,1,4]
    logger.info("Original_Array: {}".format(ar))
    result = find_next_greatest(ar)
    logger.info("Result: {}".format(result))
