from unittest import TestCase
from importlib import machinery, util

loader = machinery.SourceFileLoader('lib', './fibonacci-number-509.py')
spec = util.spec_from_loader(loader.name, loader)
lib = util.module_from_spec(spec)
loader.exec_module(lib)


class TestSolution(TestCase):
    def test_fib(self):
        cases = {
            0: 0,
            1: 1,
            2: 1,
            3: 2,
            4: 3,
            5: 5,
            6: 8,
            7: 13,
            8: 21,
            9: 34,
            10: 55,
        }

        for param, expect in cases.items():
            actual = lib.Solution.fib(param)
            self.assertEqual(actual, expect)
