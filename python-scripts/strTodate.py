import sys
import datetime

dateStr = sys.argv[1]

dateObj = datetime.datetime.strptime(dateStr,'%a %b %d %H:%M:%S %Z %Y')
print "Hour = {}".format(dateObj.hour)
print "Day = {}".format(dateObj.day)
print "Month = {}".format(dateObj.month)
print "Year = {}".format(dateObj.year)
