FROM python:3.8-slim

WORKDIR /app

COPY . /app

RUN pip install --trusted-host pypi.org -r requirements.txt

EXPOSE 8000

ENTRYPOINT ["gunicorn", "-w", "2", "-b", "0.0.0.0:8080", "runapp:app"]

#CMD ["-b", "0.0.0.0:8080", "runapp:app"]

