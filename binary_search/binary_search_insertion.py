def binary_search_insertion_simple(nums: list[int], target: int) -> int:
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
    return i


def binary_search_insertion(nums: list[int], target: int) -> int:
    i, j = 0, len(nums) - 1
    while i <= j:
        m = (i + j) // 2
        median = nums[m]
        if median >= target:
            j = m - 1
        else:
            i = m + 1

    return i


def binary_search_left_edge(nums: list[int], target: int) -> int:
    idx = binary_search_insertion(nums, target)
    if idx > len(nums) - 1 or nums[idx] != target:
        return -1
    return idx


def binary_search_right_edge(nums: list[int], target: int) -> int:
    idx = binary_search_left_edge(nums, target + 1)
    idx = idx - 1
    if nums[idx] != target or idx == -1:
        return -1
    return idx


def concat(nums, idx):
    tmp = nums[idx:]
    nums[idx] = target
    print(nums + tmp)


if __name__ == "__main__":
    nums = [1, 2, 3, 4, 6]
    target = 5
    idx = binary_search_insertion_simple(nums, target)
    concat(nums, idx)
    target = 4
    idx = binary_search_insertion_simple(nums, target)
    concat(nums, idx)

    nums = [1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 6]
    target = 5
    print(binary_search_insertion(nums, target))

    print(binary_search_left_edge(nums, target))
    print(binary_search_left_edge(nums, 7))
    print(binary_search_right_edge(nums, target))
    print(binary_search_right_edge(nums, 7))
