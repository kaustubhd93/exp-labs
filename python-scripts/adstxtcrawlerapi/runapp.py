import json
from flask import Flask, request
from getAdsTxt import get_ads_txt

app = Flask(__name__)

@app.route("/dev/get_content", methods = ["GET", "POST"])
def fetch_ads_txt():
    extPayload = json.loads(request.data)
    body = get_ads_txt(extPayload["url"])
    print("Returning response now.")
    return json.dumps({"code" : 200, "message" : body})

if __name__ == "__main__":
    app.run("0.0.0.0")
