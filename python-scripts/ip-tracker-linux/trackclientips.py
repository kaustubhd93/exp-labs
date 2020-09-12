import json
import logging
from flask import Flask, request

app = Flask(__name__)

logFileName = "/var/log/.trackclientip.log"

def py_logger(logData, logType='info', name=logFileName):
    formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')
    logger = logging.getLogger(name)

    handler = logging.FileHandler(logFileName, mode='a')
    handler.setFormatter(formatter)

    logger.setLevel(logging.INFO)
    logger.addHandler(handler)
    logger.info(logData)

    logger.handlers = []


@app.route("/logipaddr", methods = ["POST"])
def log_ip_addr():
    requestBody = json.loads(request.get_data())
    username = requestBody["username"]
    ipAddr = requestBody["ipAddr"]
    py_logger("User {} has logged in from IP address {}".format(username, ipAddr))
    return "Client IP address logged successfully."

if __name__ == "__main__":
    app.run("127.0.0.1", 9095)
