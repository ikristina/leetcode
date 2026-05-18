import p973_k_closest_points_to_origin
import pytest


@pytest.mark.parametrize(
    "points, k, expected",
    [
        ([(1, 3), (-2, 2)], 1, [(-2, 2)]),
        ([(3, 3), (5, -1), (-2, 4)], 2, [(3, 3), (-2, 4)]),
    ],
)
def test_kClosest(points, k, expected):
    assert p973_k_closest_points_to_origin.Solution().kClosest(points, k) == expected
