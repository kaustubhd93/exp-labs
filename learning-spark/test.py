from pyspark import SparkContext,SparkConf


class Custom(object):
    def split_stuff(self, text):
        words = text.split(" ")
        return len(words)
    
    # reduce(func) : Aggregate the elements of the dataset using a function func 
    # (which takes two arguments and returns one). This is from spark doc.
    # Refer url : https://spark.apache.org/docs/2.2.0/rdd-programming-guide.html#actions
    # This is from spark logs : Got job 1 (reduce at /home/centos/spark-sample/test.py:47) with 2 output partitions
    def sum_all(self, a, b):
        return a+b

def split_stuff(text):
    words = text.split(" ")
    return len(words)

def sum_all(wordList):
    return sum(wordList)

conf = SparkConf().setAppName("textcount")
sc = SparkContext(conf=conf)

# This just defines a base RDD pointing to a file, this is loaded into memory yet.
distFile = sc.textFile("samplefiles/data.txt")
# This is a tranformation. 
#trans = distFile.map(lambda stringLen:len(stringLen))
#trans.persist()
# This is an action.
# At this point spark breaks the computation into tasks to run on separate machines, and 
# each machine runs both its part of the map and a local reduction, returning only its 
# answer to the driver program.
#act = trans.reduce(lambda a,b:a+b)

print "++++++++++++++++++++++++++++++++++++++++"
#print "trans = {}".format(trans)
print "++++++++++++++++++++++++++++++++++++++++"
#print "act = {}".format(act)
print "++++++++++++++++++++++++++++++++++++++++"


#wordCount = distFile.map(split_stuff).reduce(lambda a,b: a+b)
obj = Custom()
wordCount = distFile.map(obj.split_stuff)
print "++++++++++++++++++++++++++++++++++++++++"
for ele in wordCount.collect():
    print ele
print "++++++++++++++++++++++++++++++++++++++++"
result = wordCount.reduce(obj.sum_all)
print result

