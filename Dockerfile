FROM golang:1.22.0 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]
