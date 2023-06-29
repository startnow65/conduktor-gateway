FROM golang:1.20.5-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o . ./...

FROM ubuntu:22.10
WORKDIR /app
COPY --from=builder /app/http-to-kafka ./http-to-kafka

ENTRYPOINT ["/app/http-to-kafka"]
