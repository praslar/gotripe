FROM golang:alpine

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./wait-for.sh", "mongodb:27017", "--timeout=100", "--", "./main"]

