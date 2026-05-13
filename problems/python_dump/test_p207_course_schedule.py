import p207_course_schedule


def test_can_finish():
    assert p207_course_schedule.canFinish(2, [[1, 0]]) is True
    assert p207_course_schedule.canFinish(2, [[1, 0], [0, 1]]) is False
    assert p207_course_schedule.canFinish(3, [[1, 0], [2, 1]]) is True
    assert p207_course_schedule.canFinish(3, [[1, 0], [2, 1], [0, 2]]) is False
