FROM golang:1.22 as builder

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary
RUN go build -o gbtnode main.go

# Final image
FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/gbtnode .

EXPOSE 9636

ENV RPC_URL=http://GBTNetwork:9636

CMD ["./gbtnode"]
