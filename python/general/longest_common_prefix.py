import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
https://leetcode.com/problems/longest-common-prefix/solution/
Q: Write a function to find the longest common prefix string among an array of strings. If there is no common prefix return empty string ""
Eg1: Input String:  str = ["floor", "flour", "float"]
    OutPut String: "flo"
Eg2: Input String: str = ["dog","racecar","car"] 
     OutPut String: str = ""
"""

def lcp(strs):
    """
    Logic:
    str = ["floor", "flour", "float"]
    result = "flo"
    One way to do this is by iterating through each element in  the list & find the lcp with the next element as shown below.
    For instance we will find the lcp of str[1] & str[2]
    r = lcp(str[1],str(2)) = 'flo'
    Now we will take the next element & find lcp(str[3],r).
    r = lcp(str[3],r) = lcp('float','flo') = 'flo'
    return r
    Time Complexity: O(S), where 'S' is the sum of all characters in all string in the list
    Space Complexity: O(1), there is no increment in space used along with the i/p size
    However the problem with this approach is given below.
    Assume that a very short string at the end of the list is the lcp. In this case, the above Algo will still do 'S' comparsions.
    One way to optimize this soluton is by doing vertical scanning.
    ie: We will take the first character of each element & see if they are same. If same we will check the next character & so on. We exit the loop once we find a mismatch
    Time Complexity for vertical scanning: O(S)
    Time complexity for converting a list to a set using set function is O(n), where 'n' is the number of elements in the list.
    In the worst case all elements in the list are equal & the for loop will execute for len(number of chrs in list element) times.
    Similarly set conversion will take time complexity of O(n), ie it will execute for len(strs)
    strs = ['floor', 'floor', 'floor']
    In the above example for loop execute '5' times which is len(floor).
    set function executes for 3 times inside the for loop ie len(strs)
    Now total time complexity = 3 * 5 = 15 which is equal to sum all characters in all strings
    Space Complexity Remains The Same O(1) as the space doesn't increse with the size of the i/p list
    """
    result = ""
    for item in zip(*strs):        # Execute for 3 Times [('f', 'f', 'f'), ('l', 'l', 'l')]
        if len(set(item)) == 1:    # Time Complexity O(n)
            result= result + item[0]
        else:
            break
    return result


if __name__ == "__main__":
    strs = ["floor", "flour", "float"]
    result = lcp(strs)
    logger.info(result)
