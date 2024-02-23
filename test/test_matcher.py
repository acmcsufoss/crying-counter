import unittest
from matcher import Matcher

class TestMatcher(unittest.TestCase):
    def test_exact_match(self):
        m = Matcher(["im crying", "i am crying"])
        self.assertTrue(m("im crying"))
        self.assertTrue(m("i am crying"))

    def test_partial_match(self):
        m = Matcher(["im crying", "i am crying"])
        self.assertTrue(m("im cryjing"))
        self.assertTrue(m("im cring"))

        self.assertFalse(m("im crayoning"))
        self.assertFalse(m("im cringe"))

        self.assertFalse(m("i-am-crying"))
        self.assertFalse(m("im-crying"))

    def test_ignoring_punctuation(self):
        m = Matcher(["im crying", "i am crying"])
        self.assertTrue(m("i'm crying"))
        self.assertTrue(m("i'm crying..."))
        self.assertTrue(m("~im ~crying~"))

    def test_interruption_threshold(self):
        m = Matcher(["im crying", "i am crying"])
        self.assertTrue(m("im literally crying"))
        self.assertFalse(m("i am so so so crying"))


if __name__ == "__main__":
    unittest.main()