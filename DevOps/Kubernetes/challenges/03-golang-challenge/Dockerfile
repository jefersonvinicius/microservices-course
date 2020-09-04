FROM golang as build
COPY src /app/src
RUN ls /app/src
WORKDIR /app
RUN go build -o main ./src
RUN ls .

FROM scratch
COPY --from=build /app/main /app
ENTRYPOINT ["/app/main"]