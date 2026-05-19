import p322_coin_change


def test_coin_change():
    assert p322_coin_change.coinChange([1, 2, 5], 11) == 3
    assert p322_coin_change.coinChange([2], 3) == -1
    assert p322_coin_change.coinChange([1], 0) == 0
