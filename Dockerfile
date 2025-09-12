# --- Build Stage ---
FROM golang:1.22 AS builder

WORKDIR /app

# Copy source code
COPY . .

# Build beth binary
RUN go mod init gbtnetwork || true
RUN go mod tidy
RUN make beth

# --- Runtime Stage ---
FROM debian:stable-slim

WORKDIR /root

# Copy binary + config
COPY --from=builder /app/build/bin/beth /usr/local/bin/beth
COPY genesis.json /root/genesis.json
COPY start-node.sh /root/start-node.sh

RUN chmod +x /root/start-node.sh

EXPOSE 9636

CMD ["/root/start-node.sh"]
