# Stackoverflow question : Why does the list not behave as expected ?
#  a = [[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0], [1, 1, 1]]
#  for i in a:
#       if sum(i) == 2:
#           a.remove(i)
#   print(a)
#   
#   Output: [[0, 0, 0], [0, 0, 1], [0, 1, 0], [1, 0, 0], [1, 1, 0], [1, 1, 1]]
#   Question: Element [1,1,0] should have been removed but it does not.

################################ Solution ######################################

a = [[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0], [1, 1, 1]]
# Make a static copy of the array 
b = tuple(a)
print("Input a has {} elements".format(len(a)))
print("Given list : {}".format(a))

# Iterate over the static copy
for i in b:
    print("I am this element : {} now".format(i))
    if sum(i) == 2:
        print("Removing element {}".format(i))
        a.remove(i)

print("Filtered list : {}".format(a))


# This question was later deleted by it's author voluntarily.
