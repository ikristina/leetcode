def search(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[left] <= nums[mid]:
            # left part is sorted
            if nums[left] <= target < nums[mid]:
                if nums[left] == target:
                    return left
                right = mid - 1
            else:
                if nums[right] == target:
                    return right
                if nums[mid] == target:
                    return mid
                left = mid + 1
        else:
            # right part is sorted
            if nums[mid] <= target <= nums[right]:
                if nums[right] == target:
                    return right
                if nums[mid] == target:
                    return mid
                left = mid + 1
            else:
                if nums[mid] == target:
                    return mid
                if nums[left] == target:
                    return left
                right = mid - 1
    return -1
