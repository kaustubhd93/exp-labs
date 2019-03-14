from pyspark.sql import SparkSession

spark = SparkSession.builder.appName("stuct01").getOrCreate()

df = spark.read.json("samplefiles/people.json")

df.show()
df.printSchema()
df.select("name").show()
df.filter(df.Age > 50).show()
df.groupBy("Age").count().show()

# Register the dataframe as a SQL temporary view
df.createOrReplaceTempView("people")

sqlDF = spark.sql("select * from people")
sqlDF.show()

# Temp views are only available for that particular spark
# session. Their scope ends when the spark session ends.
# Create a global temp view so that other sparksessions can
# use it.
df.createGlobalTempView("people")

spark.sql("select * from global_temp.people").show()

spark.newSession().sql("select * from global_temp.people").show()
