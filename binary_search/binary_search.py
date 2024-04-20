def binary_search(nums: list[int], target: int) -> int:
    i, j = 0, len(nums) - 1
    while i <= j:
        m = (i + j) // 2
        median = nums[m]
        if median == target:
            return median
        elif median < target:
            i = m + 1
        else:
            j = m - 1
    return -1


if __name__ == "__main__":
    nums = [10, 20, 30, 40, 50]
    target = 20
    print(binary_search(nums, target))
    target = 40
    print(binary_search(nums, target))
    target = 1
    print(binary_search(nums, target))
