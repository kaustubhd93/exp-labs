FROM python:3.7-slim

WORKDIR /app

COPY . /app

RUN pip install --trusted-host pypi.org -r requirements.txt

CMD ["python", "app.py"]

EXPOSE 8081
