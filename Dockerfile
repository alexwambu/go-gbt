FROM golang:1.22

WORKDIR /app

# Copy module files first (important for caching!)
COPY go.mod go.sum ./
RUN go mod download

# Now copy source code
COPY . .

RUN go build -o gbtnode main.go

EXPOSE 9636

CMD ["./gbtnode"]
