FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y make gcc
RUN make beth

FROM debian:stable-slim
WORKDIR /root

COPY --from=builder /app/build/bin/beth /usr/local/bin/beth
COPY genesis.json /root/genesis.json
COPY start-node.sh /root/start-node.sh

RUN chmod +x /root/start-node.sh

EXPOSE 9636

CMD ["/root/start-node.sh"]
