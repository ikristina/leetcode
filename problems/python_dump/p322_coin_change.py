from typing import List


def coinChange(coins: List[int], amount: int) -> int:
    if amount == 0:
        return 0

    # amount + 1 - "impossible so far". The worst-case valid solution uses amount coins (all 1s).
    # So anything above amount can only mean "not reachable."
    dp = [
        amount + 1 for _ in range(amount + 1)
    ]  # dp[i] - the minimum number of coins needed to make amount i.
    dp[0] = 0
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                if dp[i - coin] != amount + 1:
                    dp[i] = min(
                        dp[i], dp[i - coin] + 1
                    )  # dp[i-coin]+1 - the best way to make i - coin, plus this one coin.

    if dp[amount] > amount:
        return -1

    return dp[amount]
