import unittest

from p621_task_scheduler import leastInterval


class TestLeastInterval(unittest.TestCase):
    def test_leastInterval(self):
        tasks = ["A", "A", "A", "B", "B", "B"]
        n = 2
        expected = 8
        self.assertEqual(leastInterval(tasks, n), expected)
