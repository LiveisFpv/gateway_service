FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/main.go

CMD ["./main"]

# FROM golang:1.24

# WORKDIR /app

# COPY . .

# RUN go mod download

# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# RUN go build -gcflags="all=-N -l" -o main cmd/main.go

# CMD ["dlv", "exec", "./main", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]
