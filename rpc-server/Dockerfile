FROM golang:1.18
RUN apt-get update && apt-get install -y dos2unix
WORKDIR /app
COPY . .
RUN dos2unix ./build.sh
RUN sh ./build.sh
EXPOSE 8888
CMD ["./output/bootstrap.sh"]
