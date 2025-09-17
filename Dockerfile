FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o gbt ./cmd/gbt

EXPOSE 9636

CMD ["./gbt"]
