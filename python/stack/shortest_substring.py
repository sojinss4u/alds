import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
Q: In this problem, let's call the Alphabet string the string that consists of all uppercase English letters in the order they appear in the alphabet: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
Given a string consists of uppercase English letters, your task is to find the length of the shortest substring that contains the Alphabet string as a subsequence. 
Note:
A substring is a contiguous sequence of characters within a string. For instance "isipp" is a substring of "missisippi". "itwastimes" is a subsequence of "itwasthebestoftimes" but not a substring.
A subsequence is a sequence that can be derived from another sequence by deleting zero or more elements without changing the order of the remaining elements. For example for string "book", 
some example subsequence are "b", "ok", "oo", "bk" & "book". However "obk" & "kb" are not subsequence of "book" because they don't preserve the original order. 
Let's assume that the input string is "FORCESABCDEFDIVGHIJKLMNOPQRSTUVWXYZ". Then the smallest substring length is 29. 
"""

class Stack:
    def __init__(self):
        self.list = list()

    def push(self,data):
        self.list.append(data)

    def pop(self):
        self.list.pop()

    def top(self):
        return self.list[-1]

    def length(self):
        return len(self.list)

def shortest_substring(S):
    s = Stack()
    substring = ""
    alphabet_string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"[::-1]
    for index in range(len(alphabet_string)):
        s.push(alphabet_string[index])
    started = False
    count = 0
    for index in range(len(S)):
        stack_length = s.length()
        if stack_length == 0:
          break
        current_item = S[index]
        stack_top = s.top()
        if current_item == stack_top:
            started = True
            s.pop()
        if started:
            count+=1
            substring = substring + current_item
        logger.debug("Stack_Top: {}, Current_Element: {}, Count: {}".format(stack_top,current_item,count))
    logger.info("SubString: {}".format(substring))
    return count



if __name__ == "__main__":
    S = "DCABCDEFGHIJKLMWWWNOPQRSTUVWXYZKKKK"
    result = shortest_substring(S)
    logger.info("SubstringCount: {}".format(result))
