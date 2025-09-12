FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y make gcc
RUN make beth

FROM debian:stable-slim
WORKDIR /root

COPY --from=builder /app/build/bin/beth /usr/local/bin/beth
COPY genesis.json ./genesis.json

EXPOSE 9636

CMD ["sh", "-c", "beth --datadir /root/data init genesis.json && beth --datadir /root/data --networkid 999 --http --http.addr 0.0.0.0 --http.port 9636 --http.api eth,net,web3,personal,miner --mine --miner.threads=1"]
