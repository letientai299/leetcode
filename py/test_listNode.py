from unittest import TestCase

from list_node import ListNode


class TestListNode(TestCase):
    def test___str__(self):
        tests = [
            ([1, 2, 3, 4], "1 -> 2 -> 3 -> 4"),
            ([1], "1"),
        ]
        for test in tests:
            param = test[0]
            actual = ListNode(*param)
            if str(actual) != test[1]:
                print(f"ListNode({param}) == {actual} != {test[1]}")
                self.fail()
