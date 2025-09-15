# Use official Go image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies first
COPY go.mod ./
RUN go mod tidy

# Copy the rest of the app
COPY . .

# Build the binary
RUN go build -o gbtnode main.go

# Expose RPC port
EXPOSE 9636

# Run node
CMD ["./gbtnode"]
