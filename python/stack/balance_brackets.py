import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
1. We will push only the opening brackets to stack
2. Whenever we see a closing bracket, we will check if the top of the stack matches with closing bracket of the current item
3. If we see a closing bracket & the stack is empty then the string is imbalanced
4. Once all the brackets in the string is completed & we still have elements left in the stack, then the brackets are imbalanced
"""

class Stack:
    def __init__(self):
        self.list = list()

    def push(self,data):
        self.list.append(data)

    def pop(self):
        self.list.pop()

    def top(self):
        length = len(self.list)
        if length == 0:
            return None
        else:
            return self.list[length-1]
    def length(self):
        return len(self.list)


def check_balanced_brackets(string):
    s1 = Stack()
    opening_closing_bracket_map = {"}": "{", "]": "[", ")": "("}
    for i in range(len(string)):
        length = s1.length()
        top = s1.top()
        if string[i] in opening_closing_bracket_map.values():
            s1.push(string[i])
        elif length == 0:
            return False
        else:
            if top == opening_closing_bracket_map[string[i]]:
                s1.pop()
            else:
                return False
    if s1.length() > 0:
        return False
    else:
        return True


if __name__ == "__main__":
    string = "({}[])([)"
    balanced = check_balanced_brackets(string)
    if balanced:
        logger.info("String Is Balanced")
    else:
        logger.info("String Is Not Balanced")

