# Build stage
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o tape ./tape/main.go

# Run stage
FROM debian:stable-slim
WORKDIR /app
COPY config.yml ./
COPY secret.crt ./
COPY secret.key ./
COPY actions ./actions
COPY --from=builder /app/pkg ./pkg
COPY --from=builder /app/start.sh ./start.sh
RUN mkdir -p logs files output
EXPOSE 443
CMD ["/app/tape/main"]
