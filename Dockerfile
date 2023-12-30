FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./build/app ./cmd/main.go
CMD ["./build/app"]

FROM alpine

COPY --from=builder /app/build/app /app/app
COPY --from=builder /app/config/config.yaml /config.yaml

CMD ["/app/app"]