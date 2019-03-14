from pyspark.sql import SparkSession

spark = SparkSession.builder.appName("parquet").getOrCreate()

#df = spark.read.parquet("/user/sysadmin/tesfiles/userdata1.parquet")
df = spark.read.parquet("/user/sysadmin/tesfiles/bidderLogTest.parquet")
dfql = spark.sql("select count(distinct(u_bid_id)) from parquet.`/user/sysadmin/tesfiles/bidderLogTest.parquet`")
dfql.show()
#df.select("first_name","last_name","email").write.save("/user/sysadmin/tesfiles/fle01.parquet")

#df.write.parquet("/user/sysadmin/tesfiles/userdata1copy.parquet")
#df.show()

