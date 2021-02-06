import logging
import sys
logging.basicConfig(level=logging.INFO,stream=sys.stdout,format="%(message)s")
logger = logging.getLogger()


class Stack:
    def __init__(self):
        self.items = list()

    def is_empty(self):
        if len(self.items) == 0:
            return True
        else:
            return False

    def push(self,data):
        self.items.append(data)

    def pop(self):
        if self.is_empty():
            return False
        else:
            return self.items.pop()

    def top(self):
        index = len(self.items)
        return self.items[index-1]

    def update_top_end_time(self,data):
        index = len(self.items)
        self.items[index - 1][1] = data

    def print_stack(self):
        return self.items



if __name__ == "__main__":
    p1 = [[4,9],[1,3],[2,5],[3,4],[1,7],[5,15]]
    print("Input: {}".format(p1))
    """
    Sort intervals according to starting time. Now we can combine overlapping intervals in a linear traversal.
    In sorted array of intervals, if interval[i] doesn't overlap with interval[i-1], then interval[i+1] cannot overlap with
    interval[i-1].       
    """

    # Sort intervals based on start time [Optimized BubbleSort]

    for i in range(len(p1)-1):
        sorted = 0
        for j in range(len(p1)-1-i):
          if p1[j][0] > p1[j+1][0]:
            p1[j],p1[j+1] = p1[j+1],p1[j]
            sorted = 1
        if sorted == 0:
            break

    print("Sorted: {}".format(p1))

    s1 = Stack()
    s1.push(p1[0])
    for i in range(1,len(p1)):
        stack_top = s1.top()
        stack_top_start_time = stack_top[0]
        stack_top_end_time = stack_top[1]
        current_item = p1[i]
        current_item_start_time = current_item[0]
        current_item_end_time = current_item[1]
        if current_item_start_time < stack_top_end_time:
            if stack_top_end_time < current_item_end_time:
              s1.update_top_end_time(current_item_end_time)
        else:
            s1.push(current_item)
    merged_intervals = s1.print_stack()
    print("Merged_Output: {}".format(merged_intervals))

