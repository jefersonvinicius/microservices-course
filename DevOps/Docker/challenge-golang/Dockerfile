FROM golang:1.14.3-stretch AS builder

COPY ./main.go /go/src/
RUN go build /go/src/main.go

FROM scratch

COPY --from=builder /go/main /

ENTRYPOINT ["./main"]

