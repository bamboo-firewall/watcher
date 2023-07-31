FROM golang:alpine3.12

WORKDIR /go/src/github.com/bamboo-firewall/watcher

ADD . /go/src/github.com/bamboo-firewall/watcher

RUN go mod download

CMD ["go", "run", "cmd/main.go"]

