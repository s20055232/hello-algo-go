def bubble_sort(num):
    for j in range(len(nums) - 1, 0, -1):
        for i in range(j):
            if nums[i] > nums[i + 1]:
                nums[i + 1], nums[i] = nums[i], nums[i + 1]


def better_bubble_sort(num):
    for j in range(len(nums) - 1, 0, -1):
        flag = False
        for i in range(j):
            if nums[i] > nums[i + 1]:
                flag = True
                nums[i + 1], nums[i] = nums[i], nums[i + 1]
        if not flag:
            break


if __name__ == "__main__":
    data = [[2, 1], [3, 2, 1], [5, 3, 2, 4, 1], [1, 2, 3, 4, 5]]
    for nums in data:
        bubble_sort(nums)
        assert nums == sorted(nums), f"result is wrong {nums}"
        better_bubble_sort(nums)
        assert nums == sorted(nums), f"result is wrong {nums}"
