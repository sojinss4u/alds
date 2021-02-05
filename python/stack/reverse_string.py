import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()

"""
Reverse a given string using Stack DataStructure
"""

class Stack:
    def __init__(self):
        self.items = list()

    def push(self,data):
        self.items.append(data)

    def is_empty(self):
        if len(self.items) == 0:
            return True
        else:
            return False

    def pop(self):
        if self.is_empty():
            return False
        else:
            return self.items.pop()

if __name__ == "__main__":
    string = "telegram"
    logger.info("Input String: {}".format(string))
    s1 = Stack()
    for i in range(len(string)):
        s1.push(string[i])
    reverse_string = ""
    for _ in range(len(string)):
        status = s1.pop()
        if status:
            reverse_string = reverse_string + status
    logger.info("Output String: {}".format(reverse_string))
