# Paste this code into a new file captureIP.sh in 
# directory /etc/profile.d
if [ -z "$SSH_CLIENT" ]
then
    :
else
    clientIP=`echo $SSH_CLIENT | awk -F"=" '{print $1}' | awk '{print $1}'`
    username=`whoami`
    reqBody='{"ipAddr" : "'$clientIP'", "username": "'$username'"}'
    curl -s -o /dev/null -X POST http://127.0.0.1:9095/logipaddr -d "$reqBody"
fi
