FROM golang:1.20

RUN apt-get update && apt-get install -y vim git
RUN mkdir -p /usr/src/gosh
WORKDIR /usr/src/gosh

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/gosh ./main.go

CMD ["gosh"]