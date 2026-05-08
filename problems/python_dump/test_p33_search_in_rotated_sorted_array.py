from p33_search_in_rotated_sorted_array import search


def test_search():
    assert search([4, 5, 6, 7, 0, 1, 2], 0) == 4
    assert search([4, 5, 6, 7, 0, 1, 2], 3) == -1
    assert search([1], 0) == -1
    assert search([1], 1) == 0
    assert search([5, 1, 3], 1) == 1
    assert search([1], 1) == 0
