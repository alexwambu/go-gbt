BINARY=beth

all: build

build:
	@echo "Building $(BINARY)..."
	@go build -o $(BINARY) ./cmd/geth

clean:
	@rm -f $(BINARY)

run:
	./$(BINARY) --datadir ./data --networkid 999 --http --http.addr 0.0.0.0 --http.port 9636 --http.api "eth,net,web3,personal,miner"
