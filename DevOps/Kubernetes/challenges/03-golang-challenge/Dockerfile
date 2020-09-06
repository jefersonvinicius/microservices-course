FROM golang:alpine AS build
COPY src /app/src
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./src

FROM scratch
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]