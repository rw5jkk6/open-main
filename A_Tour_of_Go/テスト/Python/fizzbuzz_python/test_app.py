import unittest

from app import fizzbuzz

class TestFizzbuzz(unittest.TestCase):
    def testfizzbuzz(self):
        self.assertEqual(fizzbuzz(15), "fizzbuzz")
        self.assertEqual(fizzbuzz(3), "fizz")