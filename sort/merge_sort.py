def merge_sort(nums, left, right):
    if left >= right:
        return

    mid = (left + right) // 2
    merge_sort(nums, left, mid)
    merge_sort(nums, mid + 1, right)
    merge(nums, left, mid, right)


def merge(nums, left, mid, right):
    tmp = []
    left_idx, right_idx, length = left, mid + 1, 0

    while left_idx <= mid and right_idx <= right:
        if nums[left_idx] <= nums[right_idx]:
            tmp.append(nums[left_idx])
            left_idx += 1
        else:
            tmp.append(nums[right_idx])
            right_idx += 1

        length += 1

    for i in range(left_idx, mid + 1):
        tmp.append(nums[i])

    for i in range(right_idx, right + 1):
        tmp.append(nums[i])

    for i in range(len(tmp)):
        nums[left + i] = tmp[i]


if __name__ == "__main__":
    a = [2, 1, 5, 4, 3]
    merge_sort(a, 0, len(a) - 1)
    print(a)
