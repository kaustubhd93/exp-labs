import json
import requests
from flask import Flask, request

app = Flask(__name__)

@app.route("/wapi/v1/data",methods = ["POST"])
def get_data():
    req_body = request.json
    print(req_body)
    msg = {"code" : 200, "msg" : ""} 
    try:
        response = requests.post("http://core-service:8080/v1/data", json=req_body).text
        msg["msg"] = json.loads(response)
    except Exception as e:
        print(str(e))
        msg["code"], msg["msg"] = 500, "Something went wrong here..."
        return json.dumps(msg)
    return json.dumps(msg)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)

