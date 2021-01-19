FROM golang

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o main ./src

CMD ["/app/main" ]
