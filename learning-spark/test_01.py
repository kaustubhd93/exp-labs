# Below code only runs in pyspark shell 
# where in the spark context is already initialised
# when the pyspark shell starts.
from pyspark.sql.functions import *
textFile = spark.read.text("README.md")
textFile.first()
# Transformation
textFile.value.contains("spark")
# Transformation and then action
textFile.value.contains("spark").count()
# Finding the line with most no of words.
# Transformation : select(size(split(textFile.value, "\s+")).name("numWords"))
# Action : agg(max(col("numWords"))).collect() 
textFile.select(size(split(textFile.value, "\s+")).name("numWords")).agg(max(col("numWords"))).collect()
# Finding the word count of every word in the entire text.
# Transformation : select(explode(split(textFile.value, "\s+")).alias("word"))
# Action : groupBy("word").count()
wordCounts = textFile.select(explode(split(textFile.value, "\s+")).alias("word")).groupBy("word").count()
# Now display all that data
wordCounts.collect()
