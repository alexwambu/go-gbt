# Use official Go image
FROM golang:1.22 as builder

WORKDIR /app
COPY . .

# Build the binary
RUN go mod tidy && go build -o gbtnode main.go

# Runtime container
FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/gbtnode .

# Expose RPC port
EXPOSE 9636

# Run node
CMD ["./gbtnode"]
