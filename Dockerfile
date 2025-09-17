FROM golang:1.22

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the node
RUN go build -o gbt ./cmd/gbt

EXPOSE 9636

CMD ["./gbt"]

