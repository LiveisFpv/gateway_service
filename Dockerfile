FROM golang:1.24

COPY . .

RUN go mod download

WORKDIR /main

RUN go build -o main .

CMD ["./main"]