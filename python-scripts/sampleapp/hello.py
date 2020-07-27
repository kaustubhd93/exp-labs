import json
from flask import Flask

app = Flask(__name__)

@app.route("/v1/hello", methods=["GET"])
def hello():
    return json.dumps({"code":200, "message":"It worked!"})

if __name__ == "__main__":
    app.run("0.0.0.0")
