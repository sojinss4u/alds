import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

# Given an array of integers 'nums=[1,2,3,4,5]' & target=12, return the 3 indices [2,3,4] which adds up to target.

def three_sum(nums,target):
    # Time Complexity O(n^2)
    # Space Complexity O(n)
    # nums = [1,2,3,4]
    # target = 12
    for index1 in range(len(nums)):
        diff1 = target - nums[index1]
        elm_index_map = {}
        for index2 in range(index1,len(nums)):
          diff2 = diff1 - nums[index2]
          if diff2 in elm_index_map:
              return [index1,elm_index_map[diff2],index2]
          else:
              elm_index_map[nums[index2]] = index2

if __name__ == "__main__":
    nums = [1,2,3,4,5]
    target = 12
    result = three_sum(nums,target)
    logger.info(result)
