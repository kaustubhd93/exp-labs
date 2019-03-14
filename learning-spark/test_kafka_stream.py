import json
from pyspark import SparkContext,SparkConf
from pyspark.streaming import StreamingContext
from pyspark.streaming.kafka import KafkaUtils
from pyspark.sql import SparkSession,Row


conf = SparkConf().setAppName("testkafkarecvstream")
sc = SparkContext(conf=conf)
ssc = StreamingContext(sc, 300)
spark = SparkSession.builder.appName("testkafkarecvstream").getOrCreate()


def makeRows(parts):
    customRow = Row(u_bid_id=parts['u_bid_id'],
                    timestamp=parts['timestamp'])
    return customRow

def createDFToParquet(rdd):
    try:
        tableBidder = spark.createDataFrame(rdd)
        tableBidder.show()
        tableBidder.write.parquet("/user/sysadmin/tesfiles/bidderLogTest.parquet" ,mode="append")
    except ValueError:
        print "++++++++++++++++++++++++++++++"
        print "Empty Rdd"
        print "++++++++++++++++++++++++++++++"

broker = "alpha-kafka-1.vrtzads.com:2181"
topic = "bidderLog"

# Reciever method
kvs = KafkaUtils.createStream(ssc,
                              broker,
                              "raw-event-streaming-consumer-test-01",
                              {topic:5},
                              {"auto.offset.reset" : "smallest"})

lines = kvs.map(lambda x: x[1])
#counts = lines.flatMap(lambda line: line.split(" ")).map(lambda word: (word, 1)).reduceByKey(lambda a,b: a+b)
#print dir(lines)
#lines.saveAsTextFile("test")
conv = lines.map(lambda x: json.loads(x))
#table = conv.map(lambda p: Row(u_bid_id=p['u_bid_id'],timestamp=p['timestamp']))
table = conv.map(makeRows)
table.foreachRDD(createDFToParquet)
#tablebidder.show()

#print kvs
lines.pprint()
table.pprint()
#print "++++++++++++++++++++++++++++++"
#conv.pprint()
#print "++++++++++++++++++++++++++++++"

#counts.pprint()
ssc.start()
ssc.awaitTermination()

