# Script to print fibonacci series upto give number n .
# Fibonacci series : 0 1 1 2 3 5 8 
import sys

limit = int(sys.argv[1])
startingPoint = int(sys.argv[2])
fiboStore = [startingPoint, startingPoint + 1]


for i in range(limit-2):
    fiboStore.append(fiboStore[-1] + fiboStore[-2])

print(fiboStore)

