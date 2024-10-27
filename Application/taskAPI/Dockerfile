# First stage: Builder
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o taskManager main.go

# Second stage: Final image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/taskManager .

EXPOSE 8082
EXPOSE 8080
CMD ["./taskManager"]