import sys

fileName = sys.argv[1]

fObj = open(fileName, "r+")
data = fObj.readlines()
fObj.close()

#print data
sortedData = sorted(data)
for line in sortedData:
    print line

