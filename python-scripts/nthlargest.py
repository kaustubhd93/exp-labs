import sys

n = 2
unsortd = sys.argv[1:]
unsortd = map(int, unsortd)
index = len(unsortd) - 1

while index >= 0:
    for i in range(index):
        if unsortd[i] < unsortd[i+1]:
            unsortd[i], unsortd[i+1] = unsortd[i+1],unsortd[i]
    index = index - 1
    print(index)

print(unsortd)
print(unsortd[n-1])

# def bubble_sort(array):
#     index = len(array) - 1
#     while index >= 0:
#         for j in range(index):
#             if array[j] > array[j+1]:
#                 array[j], array[j+1] = array[j+1], array[j]
#         index -= 1
#     return array

# print(bubble_sort([10,100,90,6,35]))