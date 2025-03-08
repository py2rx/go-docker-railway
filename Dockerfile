# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /bin/hello ./main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest
COPY --from=builder /bin/hello /bin/hello
CMD ["/bin/hello"]