def selection_sort(nums):
    for idx in range(len(nums)):
        _min = idx
        for i in range(idx + 1, len(nums)):
            if nums[i] < nums[_min]:
                _min = i

        nums[idx], nums[_min] = nums[_min], nums[idx]


if __name__ == "__main__":
    data = [[2, 1], [3, 2, 1], [5, 3, 2, 4, 1]]
    for nums in data:
        selection_sort(nums)
        assert nums == sorted(nums), f"result is wrong {nums}"
