# need this to enable type hint using class name within the class itself
# this feature will be enabled by default in python 4, but we are still on
# python 3
from __future__ import annotations


class ListNode:
    def __init__(self, x, *argv):
        self.val = x
        self.next = None

        if not argv:
            return

        node = self
        for i in argv:
            node.next = ListNode(i)
            node = node.next

    def __str__(self):
        node = self
        result = str(node.val)
        while node.next is not None:
            node = node.next
            result += f' -> {node.val}'
        return result

    def __eq__(self, other: ListNode):
        if other is None:
            return False

        if other.val != self.val:
            return False

        if self.next is None and other.next is not None:
            return False

        return self.next.__eq__(other.next)
