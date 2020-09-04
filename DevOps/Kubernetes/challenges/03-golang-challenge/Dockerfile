FROM golang as build
COPY src /app/src
WORKDIR /app
RUN go build -o main ./src

FROM scratch
COPY --from=build /app/main /app
ENTRYPOINT ["/app/main"]