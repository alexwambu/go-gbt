#!/bin/sh
set -e

# Init chain only if not already initialized
if [ ! -d "/root/data/geth" ]; then
  echo ">>> Initializing GBTNetwork genesis..."
  beth --datadir /root/data init /root/genesis.json
fi

# Start node
echo ">>> Starting GBTNetwork node..."
exec beth --datadir /root/data \
  --networkid 999 \
  --http --http.addr 0.0.0.0 --http.port 9636 \
  --http.api eth,net,web3,personal,miner \
  --mine --miner.threads=1
