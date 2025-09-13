FROM golang:1.22

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod ./
RUN go mod tidy

# Copy everything
COPY . .

# Build beth
RUN make build

EXPOSE 9636
EXPOSE 8545
EXPOSE 10000

CMD ["./beth", "--datadir", "./data", "--networkid", "999", "--http", "--http.addr", "0.0.0.0", "--http.port", "9636", "--http.api", "eth,net,web3,personal,miner"]
