FROM golang:1.18-alpine as builder

WORKDIR /go/src/github.com/bamboo-firewall/watcher

ADD . /go/src/github.com/bamboo-firewall/watcher

RUN go mod download
RUN go build -o watcher cmd/main.go

FROM alpine:3.15

WORKDIR /app
ENV TZ=Asia/Ho_Chi_Minh
COPY --from=builder /go/src/github.com/bamboo-firewall/watcher/watcher /app/watcher
# add package for handle timezone in alpine
RUN apk add tzdata \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone
CMD [ "/app/watcher" ]


