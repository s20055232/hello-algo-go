def insertion_sort(x):
    # 這個版本也可以，但每次都會交換元素，效率比較低下
    if len(x) < 2:
        return x

    for j in range(1, len(x)):
        base = x[j]
        idx = j
        for i in range(j - 1, -1, -1):
            if base < x[i]:
                x[idx], x[i] = x[i], x[idx]
                idx -= 1

    return x


def insertion_sort_better(x):
    # 這個版本改善了反覆交換元素的浪費，先將原本的元素往右複製，直到找到位置時直接插入
    if len(x) < 2:
        return x

    for i in range(1, len(x)):
        base = x[i]
        idx = i - 1
        while base < x[i] and idx >= 0:
            x[idx + 1] = x[idx]
            idx -= 1
        x[idx + 1] = base

    return x


if __name__ == "__main__":
    use_cases = [[3, 2, 1], [5, 3, 7, 9, 1, 4], [], [2]]
    for use_case in use_cases:
        assert insertion_sort(use_case) == sorted(use_case), f"Wrong!!, {use_case}"

    for use_case in use_cases:
        assert insertion_sort_better(use_case) == sorted(
            use_case
        ), f"better Wrong!!, {use_case}"
