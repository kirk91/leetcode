# encoding: utf-8


class Iterator(object):
    def __init__(self, nums):
        """
        Initializes an iterator object to the beginning of a list.
        :type nums: List[int]
        """
        self.nums = nums
        self.pos, self.count = 0, len(nums)

    def hasNext(self):
        """
        Returns true if the iteration has more elements.
        :rtype: bool
        """
        return self.pos < self.count and True or False

    def next(self):
        """
        Returns the next element in the iteration.
        :rtype: int
        """
        if not self.hasNext():
            raise StopIteration
        val = self.nums[self.pos]
        self.pos += 1
        return val


class PeekingIterator(object):
    def __init__(self, iterator):
        """
        Initialize your data structure here.
        :type iterator: Iterator
        """
        self.iterator = iterator
        self.peeking = []

    def peek(self):
        """
        Returns the next element in the iteration without advancing the iterator.
        :rtype: int
        """
        if not self.iterator.hasNext():
            raise StopIteration
        val = self.iterator.next()
        self.peeking.append(val)
        return val

    def next(self):
        """
        :rtype: int
        """
        if self.peeking:
            return self.peeking.pop(0)
        if not self.iterator.hasNext():
            raise StopIteration
        return self.iterator.next()

    def hasNext(self):
        """
        :rtype: bool
        """
        if self.peeking:
            return True
        return self.iterator.hasNext()

if __name__ == "__main__":
    iter = PeekingIterator(Iterator([1, 2, 3, 4]))
    while iter.hasNext():
        val = iter.peek()
        next_val = iter.next()
        if next_val != val:
            raise Exception("expect %d, but got %d", val, next_val)
