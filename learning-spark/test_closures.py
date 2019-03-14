from pyspark import SparkContext,SparkConf

conf = SparkConf().setAppName("testclosures")
sc = SparkContext(conf=conf)

counter = 0
data = [1,2,3,4,5,6,7,8,9]
rdd = sc.parallelize(data)

def increment_counter(x):
    global counter
    counter += x

#result = rdd.collect()
#print "+++++++++++++++++++++++++++++++"
#print result
#print type(result)
#print "+++++++++++++++++++++++++++++++"
#
#for ele in result:
#    increment_counter(1)
    
print "Counter value : {}".format(counter)

