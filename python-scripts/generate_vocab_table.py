import sys

limit = int(sys.argv[1])
print "| **Count** | **Word** | **Synonym** | **Antonym** | **Sentence** |"
print "| ------| ------ | ------ | ------ | ------ |"
for i in range(1,limit):
    print "|{}| **** | | | |".format(i)
