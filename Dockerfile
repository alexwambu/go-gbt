# Build stage
FROM golang:1.22 AS builder
WORKDIR /app

# Copy everything including vendor
COPY . .

# Force use of vendor folder
ENV GOFLAGS=-mod=vendor

# Build binary
RUN go build -v -o build/bin/beth ./cmd/beth

# Runtime
FROM debian:bookworm-slim
WORKDIR /root
COPY --from=builder /app/build/bin/beth /usr/local/bin/beth
COPY genesis.json /root/genesis.json
RUN chmod +x /usr/local/bin/beth

EXPOSE 9636

CMD ["/usr/local/bin/beth"]
