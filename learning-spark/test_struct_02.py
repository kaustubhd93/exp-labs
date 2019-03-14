from pyspark import SparkContext,SparkConf
from pyspark.sql import Row
from pyspark.sql import SparkSession

conf = SparkConf().setAppName("struct02reflect")
sc = SparkContext(conf=conf)
spark = SparkSession.builder.appName("struct02reflect").getOrCreate()

lines = sc.textFile("samplefiles/people.txt")
parts = lines.map(lambda l: l.split(","))
#print parts.collect()
people = parts.map(lambda p: Row(Name=p[0],Age=int(p[1])))
#print people.collect()

# Infer the schema and register data frame as a table.
schemaPeople = spark.createDataFrame(people)
schemaPeople.show()

print schemaPeople.rdd.collect()
