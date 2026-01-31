# Build Stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
# Build ra file binary tên là 'main'
RUN go build -o main .

# Run Stage
FROM alpine:3.17
WORKDIR /app
# Copy file binary từ stage builder sang
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
