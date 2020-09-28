FROM golang:1.15.2-alpine as builder

COPY src/ go/src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./go/src/main.go

FROM scratch

COPY --from=builder /go/main /go/
CMD ["./go/main"]