import p198_house_robber


def test_house_robber():
    assert p198_house_robber.rob([1, 2, 3, 1]) == 4
    assert p198_house_robber.rob([2, 7, 9, 3, 1]) == 12
    assert p198_house_robber.rob([]) == 0
