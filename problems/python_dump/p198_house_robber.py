from typing import List


def rob(nums: List[int]) -> int:
    n = len(nums)
    if n == 0:
        return 0
    if n == 1:
        return nums[0]

    dp = [0 for _ in range(len(nums))]
    for i in range(len(nums)):
        total = max(dp[i - 1], nums[i] + (dp[i - 2] if i - 2 >= 0 else 0))
        dp[i] = total

    return max(dp)
