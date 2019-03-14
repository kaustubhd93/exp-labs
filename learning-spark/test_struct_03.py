from pyspark import SparkContext,SparkConf
from pyspark.sql import Row
from pyspark.sql import SparkSession
from pyspark.sql.types import *

conf = SparkConf().setAppName("struct03Programatically")
sc = SparkContext(conf=conf)
spark = SparkSession.builder.appName("struct03Programatically").getOrCreate()

lines = sc.textFile("samplefiles/people.txt")
parts = lines.map(lambda l: l.split(","))
people = parts.map(lambda p : (p[0].strip(),int(p[1].strip())))

fields = [StructField("Name",StringType(),True),StructField("Age",IntegerType(),True)]
schema = StructType(fields)

df = spark.createDataFrame(people, schema)

df.show()
