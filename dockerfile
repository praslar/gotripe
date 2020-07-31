FROM golang:alpine

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./wait-for", "mongodb:27017", "--timeout=100", "--", "./main"]

