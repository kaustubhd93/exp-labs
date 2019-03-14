from pyspark import SparkContext,SparkConf
from pyspark.sql import SparkSession
from pprint import pprint

conf = SparkConf().setAppName("testparquet")
sc = SparkContext(conf=conf)

spark = SparkSession.builder.getOrCreate()

df = spark.read.parquet("samplefiles/userdata1.parquet")
result = df.collect()
print "++++++++++++++++++++++++++++++++++++++++++"
print pprint(result)
print "++++++++++++++++++++++++++++++++++++++++++"

df.show()
