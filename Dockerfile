FROM golang:alpine

LABEL maintainer="github.com/bbt-t" description="TZ"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get github.com/BurntSushi/toml && \
    go get github.com/sirupsen/logrus && \
    go get github.com/joho/godotenv && \
    go get github.com/gorilla/mux && \
    go get github.com/lib/pq

COPY . .

RUN make

CMD ["./apiserver"]