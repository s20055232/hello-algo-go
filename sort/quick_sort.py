def median_three():
    pass


def partition(nums, left, right):
    i, j = left, right

    while i < j:
        while i < j and nums[j] >= nums[left]:
            j -= 1
        while i < j and nums[i] <= nums[left]:
            i += 1

        nums[i], nums[j] = nums[j], nums[i]

    nums[i], nums[left] = nums[left], nums[i]
    return i


def quick_sort(nums, left, right):
    if left >= right:
        return

    pivot = partition(nums, left, right)
    quick_sort(nums, left, pivot - 1)
    quick_sort(nums, pivot + 1, right)


if __name__ == "__main__":
    a = [2, 1, 5, 4, 3]
    quick_sort(a, 0, len(a) - 1)
    print(a)
