def radix_sort(array):
    max_element = max(array)
    max_digits = len(str(max_element))
    place = 1
    while max_digits > 0:
        bucket = [[] for _ in range(10)]
        for i in array:
            bucket[(i//place) % 10].append(i)
        array.clear()
        for i in bucket:
            for j in i:
                array.append(j)
        place *= 10
        max_digits -= 1
    return array
