FROM python:2.7-slim

WORKDIR /app

COPY . /app

# This will work as a single command you write in a shell script and will take in command line arguments passed while executing.
# ENTRYPOINT configures a container that will run as an executable
# References: 
# https://goinbigdata.com/docker-run-vs-cmd-vs-entrypoint/
# https://linuxconfig.org/passing-an-argument-to-the-interactive-docker-container-upon-its-launch
ENTRYPOINT ["python", "app.py"]
