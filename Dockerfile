FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o gbtnode main.go

EXPOSE 9636

CMD ["./gbtnode"]
