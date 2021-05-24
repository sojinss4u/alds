import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

""""
Q: Given an array of integers 'nums' & an integer target 'target' return the indices of two numbers,
such that they add up to target
"""

def two_sum(nums,target):
    """
    Time Complexity : O(n)
    Space Complexity : O(n)

    nums = [1,2,4,7,9]
    target = 16
    Logic: Iterate over each element in the array & check if elm' == (target - elm) is present in the hash map.
    If not present, add the current elm as key & it's index as value to hashmap. If the elm' is present in hash
    return it's value & current elm index as result.
    itr0: 16 - 1 = 15, 15 is not present in hashmap index_map = {1: 0}
    itr1: 16 - 2 = 14, 14 is present in hashmap index_map = {1: 0, 2: 1}
    itr2: 16 - 4 = 12, 12 is not present in hashmap index_map = {1: 0, 2: 1, 4: 2}
    itr3: 16 - 7 = 9,  9 is not present in hashmap index_map = {1: 0, 2: 1, 4: 2, 7: 3}
    itr4: 16 - 9 = 7,  7 is present in hashmap return(4,5) = return(current_index,index_map[7])
    """
    index_map = dict()
    for item in range(len(nums)):
        comp = target - nums[item]
        if comp in index_map:
            return [index_map[comp],item]
        else:
            index_map[nums[item]] = item
    return []


if __name__ == "__main__":
    nums = [1,2,4,7,9]
    target = 16
    result = two_sum(nums,target)
    logger.info(result)

