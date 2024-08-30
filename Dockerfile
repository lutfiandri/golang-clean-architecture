# stage 1: builder
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./dist/main ./cmd/app

# stage 2: runner
FROM alpine:latest as runner

WORKDIR /app

COPY --from=builder /app/dist/main .
COPY --from=builder /app/.env .env

CMD ["./main"]
