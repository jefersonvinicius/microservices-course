FROM golang AS build
COPY src /app/src
WORKDIR /app
RUN go build -o main ./src

FROM scratch
COPY --from=build /app /app
CMD ["/app/main"]