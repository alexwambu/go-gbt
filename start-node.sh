#!/bin/sh
set -e

DATA_DIR=/root/data

if [ ! -d "$DATA_DIR/geth" ]; then
  echo ">>> Initializing GBTNetwork genesis..."
  beth --datadir $DATA_DIR init /root/genesis.json
fi

echo ">>> Starting GBTNetwork node..."
exec beth --datadir $DATA_DIR \
  --networkid 999 \
  --http --http.addr 0.0.0.0 --http.port 9636 \
  --http.api eth,net,web3,personal,miner \
  --mine --miner.threads=1
