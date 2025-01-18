FROM golang:1.23 as builder

WORKDIR /app

COPY go.* /app/
RUN go mod download
COPY *.go /app/
RUN go build -o main .
COPY cmd/ /app/cmd/
RUN go build -o sub ./cmd/sub/sub.go
FROM gcr.io/distroless/base

COPY --from=builder /app/main /app/main
COPY --from=builder /app/sub /app/sub

CMD ["/app/main"]