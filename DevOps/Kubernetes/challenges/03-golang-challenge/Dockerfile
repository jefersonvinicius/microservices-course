FROM golang:alpine AS build
COPY src /app/src
WORKDIR /app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./src

FROM scratch
COPY --from=build /app /app
ENTRYPOINT ["/app/main"]