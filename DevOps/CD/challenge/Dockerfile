FROM golang:1.14-alpine

ENV GOPATH=/app
RUN apk update
RUN apk add git

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go get github.com/joho/godotenv

RUN go build -o main ./src

CMD ["/app/main" ]
