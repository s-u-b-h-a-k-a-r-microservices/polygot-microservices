FROM python:alpine
COPY . /app
WORKDIR /app
RUN pip install Flask==1.1.1
RUN pip install flask-restful==0.3.8
RUN pip install requests==2.22.0
ENTRYPOINT ["sh", "./run.sh"]